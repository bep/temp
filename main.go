package main

import (
	"fmt"

	avl "github.com/emirpasic/gods/trees/avltree"
)

func main() {
	tree := avl.NewWithStringComparator()

	tree.Put("A", "-a")
	tree.Put("B", "-b")
	tree.Put("1A", "-1A")
	v, _ := tree.Get("A")

	fmt.Println(v)
}
