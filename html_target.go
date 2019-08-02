package scraper

import (
	"bytes"
	"golang.org/x/net/html"
	"io"
	"strings"
)

type htmlTarget struct {
	content *html.Node
}

func (target htmlTarget) Content() *html.Node {
	return target.content
}

/*
Render pretty-renders the htmlTarget's scope
*/
func (target htmlTarget) Render() (string, error) {
	var contentWriter strings.Builder

	err := html.Render(&contentWriter, target.content)
	if err != nil {
		err = RenderingError(err)
	}

	return contentWriter.String(), err
}

func (target htmlTarget) IsValid() bool {
	return true
}

/*
NewFromNode instantiates a `Target` from an html.Node (golang.org/x/net/html).
*/
func newTargetFromNode(node *html.Node) *htmlTarget {
	concreteNode := *node
	return &htmlTarget{&concreteNode}
}

/*
NewFromBuffer instantiates a `Target` based on an http.Response (net/http).
*/
func newTargetFromBuffer(buffer io.ReadCloser) (*htmlTarget, error) {
	content, err := loadContent(buffer)

	if err != nil {
		return nil, err
	}

	return &htmlTarget{content}, nil
}

func loadContent(buffer io.ReadCloser) (node *html.Node, err error) {
	defer func() {
		// TODO: handle, or is best-effort enough?
		_ = buffer.Close()
	}()

	if !isContentValid(buffer) {
		return nil, ContentMissingError()
	}
	return html.Parse(buffer)
}

func isContentValid(content io.ReadCloser) bool {
	if written, err := io.CopyN(&bytes.Buffer{}, content, 1); written != 1 || err != nil {
		return false
	}
	return true
}
