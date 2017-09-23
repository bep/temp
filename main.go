package main

import (
	"log"
	"os"

	"io"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func main() {

	source := `package main

import (
	"fmt"
	"log"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/formatters/html"
)

func main() {
	
	var source = "foo3"

	err := quick.Highlight(os.Stdout, source, "go", "html", "monokai")

	if err != nil {
		log.Fatal(err)
	}
}`
	out, err := os.OpenFile("/Users/bep/hl.html", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	options := []html.Option{
		html.Standalone(),
		html.WithLineNumbers(),
		html.HighlightLines([][2]int{[2]int{11, 16}}),
		html.TabWidth(4)}

	formatters.Register("html", html.New(options...))

	err = highlight(out, source, "go", "html", "trac")

	if err != nil {
		log.Fatal(err)
	}
}

func highlight(w io.Writer, source, lexer, formatter, style string) error {
	l := lexers.Get(lexer)
	if l == nil {
		l = lexers.Analyse(source)
	}
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)

	f := formatters.Get(formatter)
	if f == nil {
		f = formatters.Fallback
	}

	s := styles.Get(style)
	if s == nil {
		s = styles.Fallback
	}

	s = s.Clone()
	err := s.Add(chroma.LineHighlight, "bg:#f48c42")
	if err != nil {
		panic(err)
	}

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}

	return f.Format(w, s, it)
}
