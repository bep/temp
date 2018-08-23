package main

import (
	"bytes"
	"html/template"
	"log"
	"strings"
)

func main() {
	data := map[string]interface{}{
		"style": template.HTMLAttr("background-image: url(../images/logo.png)"),
	}
	tpl := `
<section style="{{ .style }}">
</section>
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

	if !strings.Contains(result, "background-image") {
		log.Fatal(result)
	}

}
