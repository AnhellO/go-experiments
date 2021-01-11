package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routing awakes and writes to done channel")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("Main is going to call hello go routine")
	go hello(done)
	<-done
	fmt.Println("Main received data")
}
