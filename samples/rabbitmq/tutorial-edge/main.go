package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Go RabbitMQ Tutorial")

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

	queue, err := ch.QueueDeclare("TestQueue", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(queue)

	err = ch.Publish("", "TestQueue", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello World!"),
	})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Successfully connected to our RabbitMQ instance")
}
