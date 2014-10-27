package main

import (
	"fmt"
)

func main() {

	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	d := make(chan int)

	go func() {
		for {
			d <- 0
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case <-a:
			fmt.Println("aを受信した")
		case <-b:
			fmt.Println("bを受信した")
		case c <- 0:
			fmt.Println("cを受信した")
		default:
			fmt.Println("なにも受信していないです")
		}
	}
}
