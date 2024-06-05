package main

import (
	"08-events/fcutils/pkg/rabbitmq"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannelI()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "my-queue")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}

}
