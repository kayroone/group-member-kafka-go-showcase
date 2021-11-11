package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

const (
	topic         = "new-members"
	brokerAddress = "localhost:9092"
)

func main() {

	consume()
}

func consume() {

	conn, err := kafka.DialLeader(context.Background(), "tcp", brokerAddress, topic, 0)
	if err != nil {
		log.Fatal("Failed to dial leader: <", err)
	}

	err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatal("Failed to set read dead line: ", err)
	}

	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	bytes := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(bytes)
		if err != nil {
			break
		}

		member := MessageToMember(bytes[:n])
		log.Printf("Received message: Member with name '%s' added to group '%s'",
			member.Name, member.Group)
	}

	if err := batch.Close(); err != nil {
		log.Fatal("Failed to close batch: ", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("Failed to close connection: ", err)
	}
}

func MessageToMember(messageBytes []byte) Member {

	var member Member

	if err := json.Unmarshal(messageBytes, &member); err != nil {
		panic(err)
	}

	return member
}
