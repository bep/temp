package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/disintegration/gift"
	"github.com/mitchellh/hashstructure"
)

func main() {

	var slice []interface{}

	slice = append(slice, 3.0)
	slice = append(slice, 4.1)

	f1 := filter{
		Filter:  gift.Grayscale(),
		Options: 32,
	}

	f2 := filter{
		Filter:  gift.Gamma(float32(1.2)),
		Options: 32,
	}

	f3 := filter{
		Filter:  gift.Grayscale(),
		Options: 32,
	}

	var f4 gift.Filter = filter{
		Filter:  gift.Grayscale(),
		Options: 32,
	}

	fmt.Println("K", key(f1), key(f2), key(f3), key(f4))

}

func key(elements ...interface{}) string {

	var sb bytes.Buffer

	for _, element := range elements {
		hash, err := hashstructure.Hash(element, nil)
		if err != nil {
			panic(err)
		}
		sb.WriteString("_")
		sb.WriteString(strconv.FormatUint(hash, 10))
	}

	return sb.String()
}

type filter struct {
	Options interface{}
	gift.Filter
}
