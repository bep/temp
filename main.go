package main

import (
	"time"
)

func main() {
	go func() { println("Hello World!") }()

	time.Sleep(2 * time.Second)
}
