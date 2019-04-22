package targets

import (
	"golang.org/x/net/html"
	"strings"
)

type nodeTarget struct {
	name    string
	content *html.Node
}

/*
NewFromNode instantiates a `Target` from an html.Node (golang.org/x/net/html).
*/
func NewFromNode(name string, node *html.Node) *Target {
	target := Target(nodeTarget{name, node})
	return &target
}

func (target nodeTarget) Name() string {
	return target.name
}

/*
Parse returns a single html.Node instance (golang.org/x/net/html) which serves as the root of the `nodeTarget`'s hierarchical data.
Note that it re-parses the node, thus disconnecting it from its former hierarchy.
*/
func (target nodeTarget) Parse() (*html.Node, error) {
	nodes, err := html.ParseFragment(strings.NewReader(target.Content()), target.content)
	if len(nodes) != 1 {
		target.ambiguousTargetError()
	}
	return nodes[0], err
}

/*
Content pretty-renders the nodeTarget's scope
*/
func (target nodeTarget) Content() string {
	var contentWriter strings.Builder

	err := html.Render(&contentWriter, target.content)
	if err != nil {
		target.renderingError(err)
	}

	return contentWriter.String()
}
