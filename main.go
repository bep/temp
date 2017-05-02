package main

import (
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

// Main docs
func main() {
	fset := token.NewFileSet()
	basePath := "/Users/bep/go/src/github.com/spf13/hugo/tpl"

	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, fi := range files {
		if !fi.IsDir() {
			continue
		}

		packagePath := filepath.Join(basePath, fi.Name())

		d, err := parser.ParseDir(fset, packagePath, nil, parser.ParseComments)
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range d {
			p := doc.New(f, "./", 0)

			for _, t := range p.Types {
				if t.Name == "Namespace" {
					fmt.Println("  type", t.Name)
					fmt.Println("    docs:", t.Doc)
					for _, tt := range t.Methods {
						fmt.Println("       ", tt.Name)
						fmt.Println("          ", tt.Doc)
						var params []string
						for _, p := range tt.Decl.Type.Params.List {
							for _, pp := range p.Names {
								params = append(params, pp.Name)
							}
						}
						fmt.Println("          ", strings.Join(params, ", "))
					}
				}
			}
		}
	}
}
