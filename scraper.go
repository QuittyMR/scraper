/*
Package scraper provides a straightforward interface for scraping web content.
*/
package scraper

import (
	"golang.org/x/net/html"
	"io"
	"sync"
)

/*
Scraper is the base type used to scrape content.
Do not instantiate it directly - rather use one of the provided scraper.New functions
*/
type Scraper struct {
	target   Target
	lastNode *html.Node
}

/*
Filters is the input to the Scraper's Find methods. It can be populated by a tag type, parameters (see `Attributes`) or both.
Note that multiple filter arguments are resolved with an `&&` operator.

	scraperInstance.FindAll(scraper.Filters{Tag:"html"})
*/
type Filters struct {
	Tag           string
	Attributes    Attributes
	AttributeKeys AttributeKeys
	match         predicate
}

type Attributes map[string]string
type AttributeKeys []string

/*
Attributes specifies tag attributes to be searched for using the Scraper's Find methods.
It is a convenience shorthand for `map[string]string` and can contain any number of attribute sets.
Note that multiple parameters are resolved with an `&&` operator.

	scraperInstance.FindAll(scraper.Filters(Attributes:scraper.Attributes{"class":"someClass"}))
*/

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

func (scraper Scraper) GetContent() *html.Node {
	return scraper.target.Content()
}

func (scraper Scraper) GetType() string {
	return scraper.GetContent().Data
}

func (scraper Scraper) GetAttributes() Attributes {
	attributes := Attributes{}
	for _, attribute := range scraper.GetContent().Attr {
		attributes[attribute.Key] = attribute.Val
	}
	return attributes
}

func (scraper Scraper) Find(filters Filters) *Scraper {
	//TODO: Replace with a non-concurrent approach
	for result := range scraper.FindAll(filters) {
		return result
	}
	return nil
}

/*
FindAll returns all nodes matching the provided Filters
*/
func (scraper Scraper) FindAll(filters Filters) <-chan *Scraper {
	filters.build()
	operations := sync.WaitGroup{}
	matchingNodes := make(chan *Scraper, 1)

	findInNode := func(node *html.Node) {
		if filters.match(node) {
			nodeScraper, _ := NewFromNode(node)
			matchingNodes <- nodeScraper
		}
	}

	operations.Add(1)
	go searchTreeLayer(&operations, scraper.GetContent(), findInNode)

	go func() {
		operations.Wait()
		close(matchingNodes)
	}()

	return matchingNodes
}

func searchTreeLayer(operations *sync.WaitGroup, node *html.Node, callable func(node2 *html.Node)) {
	//TODO: Benchmark synchronous approach (remove goroutine calls and WaitGroup)
	for subNode := node; subNode != nil; subNode = subNode.NextSibling {
		callable(subNode)

		operations.Add(1)
		go searchTreeLayer(operations, subNode.FirstChild, callable)
	}
	operations.Done()
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
	scraper.lastNode = scraper.getLastSubNode(nil)
	return
}

func (scraper Scraper) getLastSubNode(node *html.Node) *html.Node {
	if node == nil {
		node = scraper.GetContent()
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
*/
func (filter *Filters) build() {
	if filter.match != nil {
		return
	}
	var predicates []predicate

	// Default pass-through filter
	if filter.Attributes == nil && filter.Tag == "" {
		predicates = append(predicates, func(_ *html.Node) bool { return true })
	}

	if filter.Tag != "" {
		predicateFunc := func(value string) func(node *html.Node) bool {
			return func(node *html.Node) bool {
				return node.Data == value
			}
		}(filter.Tag)

		predicates = append(predicates, predicateFunc)
	}

	for attribute, attributeValue := range filter.Attributes {
		predicateFunc := func(attribute string, attributeValue string) func(node *html.Node) bool {
			return func(node *html.Node) bool {
				for _, nodeAttribute := range node.Attr {
					if nodeAttribute.Key == attribute && nodeAttribute.Val == attributeValue {
						return true
					}
				}
				return false
			}
		}(attribute, attributeValue)

		predicates = append(predicates, predicateFunc)
	}

	for _, attributeKey := range filter.AttributeKeys {
		predicateFunc := func(attributeKey string) func(node *html.Node) bool {
			return func(node *html.Node) bool {
				for _, nodeAttribute := range node.Attr {
					if nodeAttribute.Key == attributeKey {
						return true
					}
				}
				return false
			}
		}(attributeKey)

		predicates = append(predicates, predicateFunc)
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
