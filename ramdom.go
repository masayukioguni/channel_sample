package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for {
			a <- 1
		}
	}()
	go func() {
		for {
			b <- 0
		}
	}()
	go func() {
		for {
			c <- 0
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case value := <-a:
			fmt.Println("a" + strconv.Itoa(value))
		case <-b:
			fmt.Println("b")
		case <-c:
			fmt.Println("c")
		default:
			fmt.Println("no message received")
		}
	}

}
