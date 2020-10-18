package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer Application")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	messages, err := ch.Consume("TestQueue", "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range messages {
			fmt.Printf("Received message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully connected to our RabbitMQ instance")
	fmt.Println("...waiting for messages")
	<-forever
}
