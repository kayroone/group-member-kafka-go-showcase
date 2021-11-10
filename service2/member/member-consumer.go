package main

import (
	"context"
	"fmt"
)

const (
	topic         = "new-members"
	brokerAddress = "localhost:9092"
)

func main() {

	ctx := context.Background()

	consume(ctx)
}

func consume(ctx context.Context) {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
	})

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}
