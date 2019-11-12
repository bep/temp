package main

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

func main() {
	md := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAttribute(),
		),
	)
	source := []byte(`
## heading {#id .className attrName=attrValue class="class1 class2"}
`)
	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		panic(err)
	}
}
