package main

import (
	"fmt"
	"time"
)

func doSomething(str string) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%s: %d\n", str, i)
	}
}

func main() {
	// calling this function the normal way
	doSomething("Hello")

	// Running it inside a go routine
	go doSomething("World")

	go func() {
		fmt.Println("Go routines are awesome")
	}()

	time.Sleep(2 * time.Second)
}
