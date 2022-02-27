/*
Package scraper provides a straightforward interface for scraping web content.
*/
package scraper

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"strings"
	"sync"
)

/*
Scraper is the base type used to scrape content.
Do not instantiate it directly - rather use one of the provided scraper.New functions
*/
type Scraper struct {
	target Target
}

/*
Filter is the input to the Scraper's Find methods. It can be populated by a tag type, parameters (see `Attributes`) or both.
Note that multiple filter arguments are resolved with an `&&` operator.

	scraperInstance.FindAll(scraper.Filter{Tag:"div"})
*/
type Filter struct {
	Tag        string
	Attributes Attributes
	IsExact    bool
	match      predicate
}

/*
Attributes specifies tag attributes to be searched for using the Scraper's Find methods.
It is a convenience shorthand for `map[string]string` and can contain any number of attribute sets.
Note that multiple parameters are resolved with an `&&` operator.

	scraperInstance.FindAll(scraper.Filter(Attributes:scraper.Attributes{"class":"someClass"}))
*/
type Attributes map[string]string

type predicate func(node *html.Node) bool

/*
NewFromBuffer instantiates a new Scraper instance from a given `http.Response` (net/http).
You should consider using `NewFromURI` if your requested resource is trivial to get.
Note that this function will close the `Body` handle for you.
*/
func NewFromBuffer(buffer io.ReadCloser) (*Scraper, error) {
	target, err := newTargetFromBuffer(buffer)
	if err != nil {
		return nil, err
	}
	return newFromTarget(target)
}

/*
NewFromNode instantiates a new Scraper instance from a given `html.Node` (golang.org/x/net/html).
It is used internally to allow scraping the results of a previous scrape, but provided here if you want to build a hybrid.
*/
func NewFromNode(node *html.Node) (*Scraper, error) {
	return newFromTarget(newTargetFromNode(node))
}

/*
Render returns a rendered version of the Scraper's content.
Note that the rendering is best-effort (see golang.org/x/net/html/render.go)
*/
func (scraper Scraper) Render() (string, error) {
	return scraper.target.Render()
}

/*
Content returns the node the Scraper instance is wrapping. It should be considered a lower-level API
*/
func (scraper Scraper) Content() *html.Node {
	return scraper.target.Content()
}

/*
Type returns the tag type for HTML nodes. For text nodes, it will return the text itself
*/
func (scraper Scraper) Type() string {
	return scraper.Content().Data
}

/*
Attributes returns a map of all attributes on the node
*/
func (scraper Scraper) Attributes() Attributes {
	attributes := make(Attributes)
	for _, nodeAttribute := range scraper.Content().Attr {
		attributes[nodeAttribute.Key] = nodeAttribute.Val
	}
	return attributes
}

/*
Text returns the text embedded in the node.
If other tags are nested under it, it will return an empty string and false OK
*/
func (scraper Scraper) Text() (string, bool) {
	content := scraper.Content()
	text := strings.Builder{}
	for child := content.FirstChild; child != nil; child = child.NextSibling {
		if child.Type != html.TextNode {
			return "", false
		}
		text.WriteString(child.Data)
	}
	return text.String(), true
}

/*
TextO is an optimistic version of Text that will simply return an empty string if anything goes wrong (see Text docs).
It is useful for inlining operations if you trust your inputs
*/
func (scraper Scraper) TextOptimistic() string {
	text, _ := scraper.Text()
	return text
}

/*
Find returns the first node matching the provided Filter.
Note that this method is currently very inefficient and needs to be reimplemented
*/
func (scraper Scraper) Find(filter Filter) *Scraper {
	//TODO: Replace with a non-concurrent approach
	result, ok := <-scraper.FindAll(filter)
	if !ok {
		return nil
	}
	return result
}

/*
FindAll returns all nodes matching the provided Filter
TODO: better way to track completion?
*/
func (scraper Scraper) FindAll(filter Filter) <-chan *Scraper {
	filter.build()
	matchingNodes := make(chan *Scraper)
	isMatching := func(node *html.Node) {
		if filter.match(node) {
			nodeScraper, _ := NewFromNode(node)
			matchingNodes <- nodeScraper
		}
	}

	operations := sync.WaitGroup{}
	operations.Add(1)

	go searchNode(&operations, scraper.Content(), isMatching)

	go func(operations *sync.WaitGroup) {
		operations.Wait()
		close(matchingNodes)
	}(&operations)

	return matchingNodes
}

/*
//TODO: Benchmark synchronous approach (remove goroutine calls and WaitGroup)
//TODO: can isMatching have mp side effects?
*/
func searchNode(operations *sync.WaitGroup, node *html.Node, isMatching func(node2 *html.Node)) {
	defer operations.Done()
	for subNode := node; subNode != nil; subNode = subNode.NextSibling {
		if subNode.Type == html.TextNode {
			continue
		}
		operations.Add(1)
		isMatching(subNode)

		go searchNode(operations, subNode.FirstChild, isMatching)
	}
}

/*
newFromTarget returns an instance of Scraper from any type implementing the Target interface - see targets package.
It also politely handles any errors induced by the implementation, and recovers if possible.
*/
func newFromTarget(target Target) (scraper *Scraper, err error) {
	defer func() {
		switch internalError := recover().(type) {
		case nil:
			return
		case error:
			err = internalError
		}
	}()

	if !target.IsValid() {
		target = EmptyTarget{}
	}

	scraper = &Scraper{target: target}
	return
}

func (scraper Scraper) getLastSubNode(node *html.Node) *html.Node {
	if node == nil {
		node = scraper.Content()
	}
	if node.LastChild == nil && node.NextSibling == nil {
		return node
	} else {
		return scraper.getLastSubNode(node.LastChild)
	}
}

/*
build generates a composite function that includes all filter predicates,
using closures for eager evaluation of the values.
TODO: reconsider fault tolerance
*/
func (filter *Filter) build() {
	if filter.match != nil {
		return
	}
	var predicates []predicate

	if filter.Tag != "" {
		predicateFunc := func(value string) func(node *html.Node) bool {
			return func(node *html.Node) bool {
				return node.Data == value
			}
		}(filter.Tag)

		predicates = append(predicates, predicateFunc)
	}

	for attribute := range filter.Attributes {
		predicateFunc := func(attributeKey string, attributeValue string) predicate {
			return func(node *html.Node) bool {
				for _, nodeAttribute := range node.Attr {
					NormalizedNodeAttributeValue := fmt.Sprintf(" %v ", nodeAttribute.Val)
					NormalizedAttributeValue := fmt.Sprintf(" %v ", attributeValue)
					if nodeAttribute.Key == attributeKey && strings.Contains(NormalizedNodeAttributeValue, NormalizedAttributeValue) {
						return true
					}
				}
				return false
			}
		}(attribute, filter.Attributes[attribute])

		predicates = append(predicates, predicateFunc)
	}

	// Default pass-through filter
	if len(predicates) == 0 {
		predicates = []predicate{func(_ *html.Node) bool { return true }}
	}

	filter.match = func(node *html.Node) bool {
		for _, predicate := range predicates {
			if !predicate(node) {
				return false
			}
		}
		return true
	}
}
