package member

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

const (
	topic = "new-members"
	//brokerAddress = "redpanda:29092"
	brokerAddress = "localhost:9092"
)

func Produce(newMember Member) {

	log.Println("Trying to write new member: ", newMember)

	w := &kafka.Writer{
		Addr:     kafka.TCP(brokerAddress),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer func() {
		err := w.Close()
		if err != nil {
			log.Fatal("Failed to close reader connection")
			return
		}
	}()

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(newMember.String()),
		},
	)
	if err != nil {
		log.Fatal("Failed to write messages: ", err)
	}

	log.Println("Successfully added and published new member: ", newMember)
}
