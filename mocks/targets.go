package mocks

import (
	"bytes"
)

//
//import (
//	"bitbucket.org/Quitty/scraper"
//	"bytes"
//	"golang.org/x/net/html"
//	"io/ioutil"
//	"net/http"
//	"net/url"
//	"strings"
//)
//
//func Response(content string) *http.Response {
//	return &http.Response{
//		Request: &http.Request{
//			URL: &url.URL{
//				Scheme: "http",
//				Path:   "localhost",
//			},
//		},
//		Body: ioutil.NopCloser(bytes.NewBufferString(content)),
//	}
//}
//
//func ResponseTarget(content string) scraper.Target {
//	target, _ := scraper.newTargetFromResponse(Response(content))
//	return targets.Target(target)
//}
//
//func NodeTarget(content string) targets.Target {
//	node, err := html.Parse(strings.NewReader(content))
//	if err != nil {
//		panic(err)
//	}
//	return targets.Target(targets.NewFromNode(node, "arbitrary node"))
//}

type Buffer bytes.Buffer

func (buffer Buffer) Read(p []byte) (n int, err error) {
	return buffer.Read(p)
}

func (buffer Buffer) Write(p []byte) (n int, err error) {
	return buffer.Write(p)
}

func (Buffer) Close() error {
	return nil
}
func NewBuffer() Buffer {

}
