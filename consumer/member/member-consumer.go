package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

const (
	topic         = "new-members"
	brokerAddress = "localhost:9092"
)

func main() {

	consume()
}

func consume() {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{brokerAddress},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	defer func() {
		err := r.Close()
		if err != nil {
			log.Fatal("Failed to close reader connection")
			return
		}
		log.Println("Successfully closed reader connection")
	}()

	for {
		message, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message: ", err)
			break
		}

		member := MessageToMember(message.Value)

		fmt.Printf("Received message at offset %d: %s\n", message.Offset, member)
	}

	if err := r.Close(); err != nil {
		log.Fatal("Failed to close reader: ", err)
	}
}

func MessageToMember(messageBytes []byte) Member {

	var member Member

	if err := json.Unmarshal(messageBytes, &member); err != nil {
		panic(err)
	}

	return member
}
