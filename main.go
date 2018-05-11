package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type MM map[string]interface{}

func (m MM) MyMy() string {
	return fmt.Sprintf("ABC: %v\n", m["a"])
}

type A struct {
	Func func() []string
	Map  map[string]interface{}
}

func (A) Method() []string {
	return []string{"c", "d"}
}

func main() {

	m := make(MM)

	m["a"] = "aa"
	m["b"] = "bb"

	for _, tpl := range []string{"{{ range $k, $v := .  }}-- {{ $k }}: {{ $v }}{{ end }} => {{ .MyMy }}"} {

		var buf bytes.Buffer
		tmpl, err := template.New("").Parse(tpl)
		if err != nil {
			panic(err)
		}
		if err := tmpl.Execute(&buf, m); err != nil {
			panic(err)
		}

		fmt.Println(buf.String())
	}

}
