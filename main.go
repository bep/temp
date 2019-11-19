package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark/parser"

	"github.com/yuin/goldmark"
)

func main() {
	dirname := "./goldmark_crashers"
	dir, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	names, _ := dir.Readdirnames(-1)

	for _, name := range names {
		filename := filepath.Join(dirname, name)
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Convert", name)
		convert(b)
	}
}

func convert(src []byte) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
		//extensions...,
		),
		goldmark.WithParserOptions(
			//parser.WithAttribute(),
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
		//rendererOptions...,
		),
	)

	var buf bytes.Buffer
	err := markdown.Convert([]byte(src), &buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())
}
