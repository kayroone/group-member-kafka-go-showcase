package member

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "new-members"
	brokerAddress = "localhost:9092"
)

func Produce(newMember Member, context context.Context) {

	log.Println("Trying to write new member: ", newMember)

	conn, err := kafka.DialLeader(context, "tcp", brokerAddress, topic, 0)
	if err != nil {
		log.Fatal("Failed to dial leader: ", err)
	}

	_ = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(newMember.String())},
	)
	if err != nil {
		log.Fatal("Failed to write messages: ", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("Failed to close writer: ", err)
	}

	log.Println("Successfully added and published new member: ", newMember)
}
