package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

func main() {
	var (
		tpl1 = `{{define "T1"}}T1_1{{end}}TPL1:{{template "T1"}}`
		tpl2 = `{{define "T1"}}T1_2{{end}}TPL2:{{template "T1"}}`
	)

	var buf bytes.Buffer
	tmpl := template.New("")

	tmpl1, err := tmpl.New("tpl1").Parse(tpl1)
	if err != nil {
		log.Fatal(err)
	}
	tmpl2, err := tmpl.New("tpl2").Parse(tpl2)
	if err != nil {
		log.Fatal(err)
	}

	if err := tmpl1.Execute(&buf, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())

	buf.Reset()
	if err := tmpl2.Execute(&buf, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())

}
