package main

import "fmt"

func main() {

	queue1 := make(chan string, 2)
	queue1 <- "one"
	queue1 <- "two"
	close(queue1)

	for elem := range queue1 {
		fmt.Println(elem)
	}

	fmt.Println("Done...")

	queue2 := make(chan string, 2)
	queue2 <- "one"
	queue2 <- "two"
	close(queue2)

LOOP:
	for {
		select {
		case elem, ok := <-queue2:
			if !ok {
				break LOOP
			}
			fmt.Println(elem)
		}
	}

	fmt.Println("Done...")

}
