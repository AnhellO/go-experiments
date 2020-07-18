package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func main() {
	fmt.Println("I'm listening")

	c := fanIn(boring("Ann"), boring("Joe"))
	for i := 0; i < 10; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}

	fmt.Println("You're both boring; I'm leaving")
}

func boring(msg string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool) // Shared between all messages.

	go func() {
		for i := 0; ; i++ {
			c <- Message{
				fmt.Sprintf("%s %d", msg, i),
				waitForIt,
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)

	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}
