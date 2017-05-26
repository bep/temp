package main

import (
	"bytes"
	"fmt"
	"html/template"
)

type A struct {
	SliceFunc func() []string
}

func (A) SliceMethod() []string {
	return []string{"c", "d"}
}

func main() {
	a := A{}
	a.SliceFunc = a.SliceMethod

	for _, v := range a.SliceMethod() {
		fmt.Println(">>", v)
	}

	for _, v := range a.SliceFunc() {
		fmt.Println(">>", v)
	}

	for _, tpl := range []string{"{{ range .SliceMethod  }}{{ . }}{{ end }}", "{{ range .SliceFunc  }}{{ . }}{{ end }}"} {

		var buf bytes.Buffer
		tmpl, err := template.New("").Parse(tpl)
		if err != nil {
			panic(err)
		}
		if err := tmpl.Execute(&buf, a); err != nil {
			panic(err)
		}

		fmt.Println(buf.String())
	}

}
