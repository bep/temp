package main

import (
	"bytes"
	"fmt"
	"log"
	"runtime/debug"

	"github.com/yuin/goldmark/parser"

	"github.com/yuin/goldmark"
)

func main() {

	convert(`#
# FOO`)
}

func convert(src string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic:\n", string(debug.Stack()))
		}
	}()

	markdown := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)
	var buf bytes.Buffer
	err := markdown.Convert([]byte(src), &buf)
	if err != nil {
		log.Fatal(err)
	}
}
