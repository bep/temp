package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// https://github.com/golang/go/blob/master/src/text/template/exec.go#L402
// https://github.com/golang/go/blob/master/src/text/template/template.go#L222

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	var clonesMu sync.Mutex
	clones := make(map[string]*template.Template)

	var wg sync.WaitGroup

	words := make([]string, 22)
	for i := 0; i < 22; i++ {
		words[i] = fmt.Sprintf("<word%d>", i)
	}

	scripts := make([]string, 22)
	for i := 0; i < 22; i++ {
		scripts[i] = js
	}

	data := map[string]interface{}{
		"Words":   words,
		"Scripts": scripts,
	}

	templateTemplate := `
<h1>This is a template</h1>
{{ range .Words }}
<a href="{{ . }}">
{{ end }}
{{ range .Scripts }}
{{ . }}
{{ end }}
{{ template "inline" . }}
{{ define "inline" }}
{{ $first := index .Words 0 }}
<a href="{{ $first }}">
{{ end }}
`

	tpl := template.New("")
	for i := 0; i < 100; i++ {
		_, err := tpl.New(fmt.Sprintf("innert%d", i)).Parse(templateTemplate)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < 20; i++ {
		var templ strings.Builder
		for j := 0; j < 60; j++ {
			templ.WriteString(fmt.Sprintf(`{{ template "innert%d" . }}`, rnd.Intn(100)))
		}
		templ.WriteString(templateTemplate)
		name := fmt.Sprintf("outert%d", i)
		_, err := tpl.New(name).Parse(templ.String())
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < 20; i++ {
		name := fmt.Sprintf("outert%d", i)
		t := tpl.Lookup(name)
		if t == nil {
			log.Fatal("not found")
		}
		tc, err := t.Clone()
		if err != nil {
			log.Fatal(err)
		}
		clones[name] = tc

	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				name := fmt.Sprintf("outert%d", rand.Intn(20))
				var templ *template.Template
				r := rand.Intn(10)
				if r < 5 {
					clonesMu.Lock()
					templ = clones[name]
					clonesMu.Unlock()
				} else if r < 7 {
					templates := tpl.Templates()
					templ = templates[rnd.Intn(len(templates))]
				} else {
					templ = tpl.Lookup(name)
				}

				if templ == nil {
					log.Fatalf("%q not found", name)
				}
				var buf bytes.Buffer
				if err := templ.Execute(&buf, data); err != nil {
					log.Fatal(err)
				}
				if buf.Len() == 0 {
					log.Fatal("nothing produced")
				}
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("Done...")
}

const js = `
<!-- Google Tag Manager -->
(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
})(window,document,'script','dataLayer','GTM-XXXX');
<!-- End Google Tag Manager -->
`
