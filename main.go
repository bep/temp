package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type Data struct {
	counter uint32
}

func (d *Data) Incr() uint32 {
	atomic.AddUint32(&d.counter, 1)
	return atomic.LoadUint32(&d.counter)
}

func main() {
	templ, err := template.New("master").Parse(`
{{ .Incr }}
{{ . }}
`)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	data := &Data{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := 0; i < 10; i++ {
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
