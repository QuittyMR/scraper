package scraper

import "golang.org/x/net/html"

/*
Target represents a scope that can be parsed into structured data or rendered as such.
It is an implementation detail meant to allow better encapsulation for the different ways of instantiating a Scraper.
*/
type Target interface {
	// Render returns a pretty-rendered version of the target's scope
	Render() (string, error)
	// Render returns the tree-structure representation of the target
	Content() *html.Node
	IsValid() bool
}
