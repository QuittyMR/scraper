package targets

import "golang.org/x/net/html"

/*
Target represents a scope that can be parsed into structured data or rendered as such.
It is an implementation detail meant to allow better encapsulation for the different ways of instantiating a Scraper.
*/
type Target interface {
	// Name returns the human-readable identifier of the target
	Name() string
	// Content returns a pretty-rendered version of the target's scope
	Content() string
	// Parse generates a tree structure out of the target's scope
	Parse() (*html.Node, error)
}
