package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

var (
	overlayTmpl *template.Template
	masterTmpl  *template.Template
)

func main() {

	templ, err := template.New("master").Parse(`
<!DOCTYPE HTML>
<html>
<head>
<title>
	{{ .title }}
</title>
</head>
<body>
</body>
</html>
`)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := 0; i < 1000; i++ {
				data := map[string]interface{}{
					"title": fmt.Sprintf("Title %d", i),
				}
				if err := templ.Execute(ioutil.Discard, data); err != nil {
					log.Fatal(err)
				}
				time.Sleep(23 * time.Millisecond)
			}
		}()
	}

	wg.Wait()
	log.Println("Done ...")
}
