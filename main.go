package main

import (
	"fmt"
	"strings"

	radix "github.com/armon/go-radix"
)

//

func main() {

	t := radix.New()
	sects := radix.New()

	sects.Insert("/blog/sect2", "sect")
	sects.Insert("/blog/sect2/s2_1", "sect")

	t.Insert("/blog/sect2/__b_h__index.md", "_index.md")
	t.Insert("/blog/sect2/foo/__b_h__index.md", "_index.md")

	t.Insert("/blog/sect2/foo/bar", "_index.md")
	t.Insert("/blog/sect2/b1/__b_h_index.md", "index.md")
	t.Insert("/blog/sect2/b1/__b_i_data.json", "data.json")
	t.Insert("/blog/sect2/b1/__b_i_b2/sunset.jpg", "sunset.jpg")

	printSection(t, "/blog/sect2/")
}

func printSection(t *radix.Tree, prefix string) {
	level := strings.Count(prefix, "/")
	t.WalkPrefix(prefix, func(s string, v interface{}) bool {
		currentLevel := strings.Count(s, "/")
		if currentLevel != level {
			return true
		}

		fmt.Println(s)

		return false
	})
}
