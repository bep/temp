package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strings"
)

func main() {
	data := map[string]interface{}{
		"abc": []string{"a", "b", "c"},
	}
	tpl := `
{{ range $i, .abc }}
{{ $i }}
{{ end }}
`

	var buf bytes.Buffer
	tmpl, err := template.New("").Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}

	if err := tmpl.Execute(&buf, data); err != nil {
		log.Fatal(err)
	}

	result := strings.TrimSpace(buf.String())

	fmt.Println(result)

}
