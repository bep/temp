package main

import (
	"fmt"
	"path"
	"strings"

	radix "github.com/armon/go-radix"
)

//

func main() {

	t := radix.New()

	t.Insert("/blog__hb_/a__hl_", "b")
	t.Insert("/blog__hb_/b/c__hl_", "c")

	getBundle := func(s string) string {

		p := path.Dir(strings.TrimPrefix(s, "/blog"))

		parts := strings.Split(p, "/")[1:]

		for i := len(parts); i >= 0; i-- {
			key := "/blog__hb_/" + strings.Join(parts[:i], "/") + "__hl_"
			prefix, _, _ := t.LongestPrefix(key)

			if prefix != "" {
				return prefix
			}

		}

		return ""

	}

	//

	fmt.Println("S:", getBundle("/blog/b/c/data.json"))
}
