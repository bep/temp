package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	b := new(bool)
	c := false

	v := viper.New()

	v.Set("a1", b)
	v.Set("a2", c)

	fmt.Println(">>>", v.GetBool("a1"))
	fmt.Println(">>>", v.GetBool("a2"))

	*b = true
	c = true

	fmt.Println(">>>", v.GetBool("a1"))
	fmt.Println(">>>", v.GetBool("a2"))

}
