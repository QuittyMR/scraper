/*
Package scraper provides a straightforward interface for scraping web content.
*/
package scraper

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"quitty.tech/scraper/targets"
)

/*
Scraper is the base type used to scrape content.
Do not instantiate it directly - rather use one of the provided scraper.New functions
*/
type Scraper struct {
	target *targets.Target
	Scope  *html.Node
}

/*
Filters is the input to the Scraper's Find methods. It can be populated by a tag type, parameters (see `Parameters`) or both.
Note that multiple filter arguments are resolved with an `&&` operator.

	scraperInstance.FindAll(scraper.Filters{Tag:"html"})
*/
type Filters struct {
	Tag        string
	Parameters Parameters
	match      predicate
}

/*
Parameters specifies tag attributes to be searched for using the Scraper's Find methods.
It is a convenience shorthand for `map[string]string` and can contain any number of attribute sets.
Note that multiple parameters are resolved with an `&&` operator.

	scraperInstance.FindAll(scraper.Filters(Parameters:scraper.Parameters{"class":"someClass"}))
*/
type Parameters map[string]string

type predicate func(node *html.Node) bool

func (scraper Scraper) String() string {
	return fmt.Sprintf("Analysis of %v", (*scraper.target).Name())
}

/*
NewFromURI instantiates a new Scraper instance from a given uri string.
It is a convenience method that handles the mandatory (net/http) boilerplate.
*/
func NewFromURI(uri string) (*Scraper, error) {
	return newFromTarget(targets.NewFromURI(uri))
}

/*
NewFromResponse instantiates a new Scraper instance from a given `http.Response` (net/http).
You should consider using `NewFromURI` if your requested resource is trivial to get.
Note that this function will close the `Body` handle for you.
*/
func NewFromResponse(response *http.Response) (*Scraper, error) {
	return newFromTarget(targets.NewFromResponse(response))
}

/*
NewFromNode instantiates a new Scraper instance from a given `html.Node` (golang.org/x/net/html).
It is used internally to allow scraping the results of a previous scrape, but provided here if you want to build a hybrid.
*/
func NewFromNode(node *html.Node) (*Scraper, error) {
	return newFromTarget(targets.NewFromNode(node.Data, node))
}

/*
Content returns a rendered version of the Scraper's content.
Note that the rendering is best-effort (see golang.org/x/net/html/render.go)
*/
func (scraper Scraper) Content() string {
	return (*scraper.target).Content()
}

/*
Find returns the first node encountered matching the provided Filters (depth-first traversal)
*/
func (scraper Scraper) Find(filter Filters) *Scraper {
	//TODO: Find elegant way to merge with FindAll
	filter.build()
	incomingNodes := make(chan *html.Node, 1)
	go func() {
		findAll(incomingNodes, scraper.Scope, filter, false)
	}()

	for matchingNode := range incomingNodes {
		nodeName := fmt.Sprintf("%v.%v", (*scraper.target).Name(), matchingNode.Data)
		nodeScraper, _ := newFromTarget(targets.NewFromNode(nodeName, matchingNode))
		return nodeScraper
	}
	return nil
}

/*
FindAll returns all nodes matching the provided Filters
*/
func (scraper Scraper) FindAll(filter Filters) (matchingNodes []*Scraper) {
	filter.build()
	incomingNodes := make(chan *html.Node, 1)
	go func() {
		findAll(incomingNodes, scraper.Scope, filter, true)
		close(incomingNodes)
	}()

	for matchingNode := range incomingNodes {
		nodeName := fmt.Sprintf("%v.%v", (*scraper.target).Name(), matchingNode.Data)
		nodeScraper, _ := newFromTarget(targets.NewFromNode(nodeName, matchingNode))
		matchingNodes = append(matchingNodes, nodeScraper)
	}
	return
}

func findAll(matchingNodes chan<- *html.Node, node *html.Node, filter Filters, isExhaustive bool) {
	for subNode := node; subNode != nil; subNode = subNode.NextSibling {
		if filter.match(subNode) {
			matchingNodes <- subNode
			if !isExhaustive {
				close(matchingNodes)
				return
			}
		}
		if subNode.FirstChild != nil {
			findAll(matchingNodes, subNode.FirstChild, filter, isExhaustive)
		}
	}
}

/*
newFromTarget returns an instance of Scraper from any type implementing the Target interface - see targets package.
It also politely handles any errors induced by the implementation, and recovers if possible.
*/
func newFromTarget(target *targets.Target) (scraper *Scraper, err error) {
	defer func() {
		switch internalError := recover().(type) {
		case nil:
			return
		case error:
			err = internalError
		}
	}()

	scraper = &Scraper{target: target}
	node, err := (*scraper.target).Parse()
	scraper.Scope = node
	return
}

/*
build generates a composite function that includes all filter predicates,
using closures for eager evaluation of the values.
*/
func (filter *Filters) build() {
	var predicates []predicate

	// Default noFilter
	if filter.Parameters == nil && filter.Tag == "" {
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

	for key, value := range filter.Parameters {
		predicateFunc := func(value string) func(node *html.Node) bool {
			return func(node *html.Node) bool {
				for _, attribute := range node.Attr {
					if attribute.Key == key && attribute.Val == value {
						return true
					}
				}
				return false
			}
		}(value)
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
