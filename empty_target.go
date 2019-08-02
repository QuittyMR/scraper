package scraper

import "golang.org/x/net/html"

type EmptyTarget struct {
	name    string
	content *html.Node
}

var emptyTarget = EmptyTarget{name: "empty target", content: &html.Node{}}

func (EmptyTarget) Render() (string, error) {
	return "", emptyTarget.RenderingError()
}

func (EmptyTarget) Content() *html.Node {
	return emptyTarget.content
}

func (EmptyTarget) IsValid() bool {
	return false
}
