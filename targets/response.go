package targets

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
)

type responseTarget struct {
	uri      string
	response *http.Response
	content  string
}

/*
NewFromResponse instantiates a `Target` based on an http.Response (net/http).
*/
func NewFromResponse(response *http.Response) *Target {
	target := Target(responseTarget{
		uri:      response.Request.URL.String(),
		response: response,
	})

	return &target
}

/*
NewFromURI instantiates a `Target` based on a uri string - which means it resolves the http.response (net/http) on its own.
*/
func NewFromURI(uri string) *Target {
	responseTarget := responseTarget{uri: uri}
	responseTarget.loadResponse()
	target := Target(responseTarget)
	return &target
}

func (target responseTarget) Name() string {
	if target.response != nil {
		return fmt.Sprintf("%v[%v]", target.uri, target.response.StatusCode)
	}
	return target.uri

}

/*
Content returns the contents of the `responseTarget`'s body.
Note that this method is stateful; calling it for the first time will store the results in the responseTarget.
Subsequent calls will retrieve the stored contents
*/
func (target responseTarget) Content() string {
	if target.content == "" {
		target.loadContent()
	}
	return target.content
}

/*
Parse returns a single html.Node instance (golang.org/x/net/html) which serves as the root of the `responseTarget`'s hierarchical data.
Note that it will error out if the `responseTarget.response.Body` object has been closed - this will be fixed in a future version.
*/
func (target responseTarget) Parse() (node *html.Node, err error) {
	return html.Parse(target.response.Body)
}

func (target *responseTarget) loadResponse() {
	response, err := http.Get(target.uri)

	if err != nil {
		target.UriError(err)
	} else if response.Body == nil {
		target.ContentMissingError()
	}
	target.response = response
}

func (target *responseTarget) loadContent() {
	bytesRead, err := ioutil.ReadAll(target.response.Body)
	if err != nil {
		target.marshallingError(err)
	}

	target.content = string(bytesRead)

	err = target.response.Body.Close()
	if err != nil {
		target.bodyHandleError(err)
	}
}
