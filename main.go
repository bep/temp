package main

import (
	"fmt"
	"reflect"
	"runtime"
)

type A struct {
}

func (a *A) M() {

}

func main() {
	a := &A{}

	name := runtime.FuncForPC(reflect.ValueOf(a.M).Pointer()).Name()
	fmt.Println("Name:", name)
}
