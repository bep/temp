package main

import (
	"log"
	"math/big"
	"time"

	"github.com/bep/typedmapcodec"

	"github.com/kr/pretty"
)

func main() {
	mi := map[string]interface{}{
		"vstring":   "Hello",
		"vint":      32,
		"vrat":      big.NewRat(1, 2),
		"vtime":     time.Now(),
		"vduration": 3 * time.Second,
		"vsliceint": []int{1, 3, 4},
		"nested": map[string]interface{}{
			"vint":      55,
			"vduration": 5 * time.Second,
		},
		"nested-typed-int": map[string]int{
			"vint": 42,
		},
		"nested-typed-duration": map[string]time.Duration{
			"v1": 5 * time.Second,
			"v2": 10 * time.Second,
		},
	}

	c, err := typedmapcodec.New()
	if err != nil {
		log.Fatal(err)
	}

	data, err := c.Marshal(mi)
	if err != nil {
		log.Fatal(err)
	}
	m := make(map[string]interface{})
	if err := c.Unmarshal(data, &m); err != nil {
		log.Fatal(err)
	}

	pretty.Print(m)

}
