package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("I'm listening")

	ann := boring("boring Ann!")
	joe := boring("boring Joe!")
	for i := 0; i < 10; i++ {
		fmt.Println(<-ann)
		fmt.Println(<-joe)
	}

	fmt.Println("You're boring; I'm leaving")
}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}
