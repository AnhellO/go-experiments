package main

import (
	"fmt"
)

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
}

func main() {
	ch := make(chan int)
	go producer(ch)

	for v := range ch {
		fmt.Printf("Received %d\n", v)
	}
}
