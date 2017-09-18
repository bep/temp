package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

func main() {

	yml := `
- title: "Harry Potter and the Order of the Phoenix"
  author: "J.K. Rowlings"
  pages: 870
- title: "A Wrinkle in Time"
  author: "Madeleine L'Engle"
  pages: 240
- title: "My Side of the Mountain"
  author: "Jean Craighead George"
  pages: 175
`

	var out interface{}

	err := yaml.Unmarshal([]byte(yml), &out)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Out: %T:%v", out, out)
}
