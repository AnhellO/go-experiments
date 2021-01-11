package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 100
}

func main() {
	ch := make(chan int)
	go sendData(ch)
	fmt.Println(<-ch)
}
