package main

import "08-events/fcutils/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannelI()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello, World!", "amq.direct")
}
