package main

import (
	"fmt"
	"strconv"
)

func main() {
	channel := make(chan int)
	go func(s chan<- int) {
		for i := 0; i < 5; i++ {
			fmt.Println("s <-" + strconv.Itoa(i))
			s <- i

		}
		close(s)
	}(channel)

	for {
		value, ok := <-channel //チャネルを受信、valueには値が、okにはクローズしたか判別できるbool型が入る
		if !ok {
			fmt.Println("channel:" + strconv.FormatBool(ok))
			break
		}
		fmt.Println(strconv.Itoa(value) + " <- channel:" + strconv.FormatBool(ok))
	}
}
