package member

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

const (
	topic         = "new-members"
	brokerAddress = "localhost:9092"
)

func Produce(newMember Member) {

	log.Println("Trying to write new member: ", newMember)

	w := &kafka.Writer{
		Addr:     kafka.TCP(brokerAddress),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(newMember.String()),
		},
	)
	if err != nil {
		log.Fatal("Failed to write messages: ", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("Failed to close writer: ", err)
	}

	log.Println("Successfully added and published new member: ", newMember)
}
