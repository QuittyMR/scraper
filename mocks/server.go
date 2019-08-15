package mocks

import (
	"bytes"
	"fmt"
	"net/http"
)

var Snippets = map[string]string{
	"span": `
		<span class="testClass1">
			Test text 1
		</span>
	`,
	"text input": `
		<input type="text"/>
	`,
	"link": `
		<a href="http://localhost">
			Test link 1
		</a>
	`,
}

var HTMLSnippet = fmt.Sprintf(`
	<html>
		<head>
		</head>
		<body>
			<div>
				%v
			</div>
			<div>
				%v
				%v
				%v
			</div>
		</body>
	</html>
`,
	Snippets["span"], Snippets["text input"], Snippets["text input"], Snippets["link"],
)

type mockHandler struct{}

func (mockHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	_, _ = writer.Write(bytes.NewBufferString(HTMLSnippet).Bytes())
	fmt.Println("Replied to web request")
}

//var MockServer = httptest.NewUnstartedServer(mockHandler{})
