package member

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "new-members"
	brokerAddress = "localhost:9092"
)

func Produce(newMember Member, ctx context.Context) {

	fmt.Println("Trying to write new member: ", newMember)

	logger := log.New(os.Stdout, "Kafka writer: ", 0)

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  logger,
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key:   nil,
		Value: []byte(newMember.String()),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}

	fmt.Println("Successfully added and published new member: ", newMember)
}
