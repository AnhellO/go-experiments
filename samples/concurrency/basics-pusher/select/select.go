package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		chan1 <- 1
		time.Sleep(5 * time.Second)
	}()

	go func() {
		chan2 <- 2
		time.Sleep(5 * time.Second)
	}()

	for i := 0; i < 2; i++ {
		select {
		case v := <-chan1:
			fmt.Printf("Value: %d\n", v)
		case v := <-chan2:
			fmt.Printf("Value: %d\n", v)
		}
	}
}
