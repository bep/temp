package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"
)

func main() {
	var buf bytes.Buffer
	tmpl := template.New("")
	setFuncs(tmpl, "init")

	var err error
	tmpl, err = tmpl.Parse(`{{ hello}}`)
	if err != nil {
		log.Fatal("parse failed:", err)
	}

	setFuncs(tmpl, "changed1")

	if err := tmpl.Execute(&buf, nil); err != nil {
		log.Fatal(err)
	}

	result := strings.TrimSpace(buf.String())
	fmt.Println(result)

}

func setFuncs(templ *template.Template, name string) {

	funcs := template.FuncMap{
		"hello": func() string { return name },
	}

	templ.Funcs(funcs)

}
