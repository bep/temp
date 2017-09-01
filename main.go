package main

import (
	"fmt"
	"sync"
)

type cache struct {
	mm sync.Map
}

func (c *cache) get(key interface{}, fn func() interface{}) interface{} {
	if v, ok := c.mm.Load(key); ok {
		return v
	}

	if v, loaded := c.mm.LoadOrStore(key, fn()); loaded {
		return v
	} else {
		return v
	}

}

func main() {

	fn := func() interface{} {
		fmt.Println("fn executed")
		return "val"
	}

	c := &cache{}

	fmt.Println(c.get("key", fn))
	fmt.Println(c.get("key", fn))
}
