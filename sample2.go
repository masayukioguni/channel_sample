package main

import (
	"fmt"
	"os"
	"strconv"
)

const goroutines = 10

func main() {
	counter := make(chan int, 1)
	for i := 0; i < goroutines; i++ {
		go func(counter chan int) {
			fmt.Println("value := <-counter")
			value := <-counter
			value++
			if value == goroutines {
				fmt.Println("exit")

				os.Exit(0)
			}

			fmt.Println("counter:= <-" + strconv.Itoa(value))
			counter <- value
		}(counter)

	}
	fmt.Println("counter <- 0 s")
	counter <- 0
	fmt.Println("counter <- 0 e")
	for {

	}

}
