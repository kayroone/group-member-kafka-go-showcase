package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

const (
	topic = "new-members"
	//brokerAddress = "redpanda:29092"
	brokerAddress = "localhost:9092"
)

func main() {

	go consumeMain()
	consumeDeadLetter()
}

func consumeMain() {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{brokerAddress},
		Topic:     topic,
		GroupID:   "member-group-1",
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
		message, err := r.FetchMessage(context.Background())
		if err != nil {
			log.Println("Error reading message: ", err)
			break
		}

		member, err := MessageToMember(message.Value)
		if err != nil {
			log.Printf("Failed to unmarshal message: %s\n", err)
			go sendToDeadLetterTopic(message)
		}

		if err := r.CommitMessages(context.Background(), message); err != nil {
			log.Printf("Failed to commit message %s with error %s\n", message.Value, err)
			break
		}

		if err == nil {
			log.Printf("Received message at offset %d: %s\n", message.Offset, member)
		}
	}
}

func MessageToMember(messageBytes []byte) (Member, error) {

	var member Member
	err := json.Unmarshal(messageBytes, &member)

	return member, err
}

func sendToDeadLetterTopic(message kafka.Message) {

	log.Printf("Sending invalid message to dead letter: '%s'\n", message.Value)

	w := &kafka.Writer{
		Addr:     kafka.TCP(brokerAddress),
		Balancer: &kafka.LeastBytes{},
	}
	defer func() {
		err := w.Close()
		if err != nil {
			log.Fatal("Failed to close ded letter writer connection")
			return
		}
		log.Println("Successfully closed dead letter writer connection")
	}()

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Value: message.Value,
			Topic: "dead-letter",
		},
	)
	if err != nil {
		log.Fatal("Failed to write messages: ", err)
		return
	}

	log.Printf("Invalid message transferred to dead letter topic: '%s'\n", message.Value)
}

func consumeDeadLetter() {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{brokerAddress},
		Topic:     "dead-letter",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	defer func() {
		err := r.Close()
		if err != nil {
			log.Fatal("Failed to close dead letter reader connection")
			return
		}
	}()

	for {
		message, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message: ", err)
			break
		}

		log.Printf("Received invalid message in dead letter topic: '%s'\n", message.Value)
	}
}
