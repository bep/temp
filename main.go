package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

type Niller interface {
	Nil() Niller
	Foo() string
}

type Nill struct {
	nilv Niller
}

func (n Nill) Nil() Niller {
	return n.nilv
}

func (n Nill) Foo() string {
	return "asdf"
}

func main() {

	tmpl, err := template.New("").Parse(`{{ with .Nil }}Failed, got {{ . }}{{ else }}OK{{ end }}`)
	if err != nil {
		log.Fatal(err)
	}

	var (
		nil1 Niller = (*Nill)(nil)
		nil2 Niller
		nil3 *Nill
		nil4 Niller = nil
	)

	for i, niller := range []Niller{nil1, nil2, nil3, nil4} {

		var buff bytes.Buffer
		n := &Nill{nilv: niller}
		a, b := template.IsTrue(n.Nil())
		err = tmpl.Execute(&buff, n)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println(i+1, a, b, buff.String())
	}
}
