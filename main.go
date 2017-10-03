package main

import (
	"fmt"
	"reflect"
)

type Helloer interface {
	Hi()
}

type (
	Hi1      int
	Hi2      int
	Helloers []Helloer
)

func (h Hi1) Hi() {
	fmt.Println("Hi", h)
}

func (h Hi2) Hi() {
	fmt.Println("Hi2", h)
}

func (h Helloers) Filter(by Helloer) Helloers {
	var filtered Helloers

	tp := reflect.TypeOf(by)

	for _, helloer := range h {
		htp := reflect.TypeOf(helloer)
		if htp.AssignableTo(tp) {
			filtered = append(filtered, helloer)
		}

	}
	return filtered
}

func main() {
	h1 := Hi1(1)
	h11 := Hi1(11)
	h2 := Hi2(2)
	hellos := Helloers{h1, h2, h1, h11}

	filtered := hellos.Filter((Hi1)(0))

	fmt.Println("Filtered:", filtered)
	// => Filtered: [1 1 11]

}
