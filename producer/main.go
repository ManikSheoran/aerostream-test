package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	brokerAddr := "localhost:9092"
	topic := "test-topic"

	// Create topic if it doesn't exist (AeroStream doesn't auto-create)
	conn, err := kafka.Dial("tcp", brokerAddr)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	conn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})
	conn.Close()

	writer := &kafka.Writer{
		Addr:         kafka.TCP(brokerAddr),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond,
	}
	defer writer.Close()

	for i := 0; i <= 10; i++ {
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("Key %v", i)),
			Value: []byte(fmt.Sprintf("%d. Message: %v", i, time.Now().Format(time.RFC1123))),
		}

		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Fatalf("failed to write message %d: %v", i, err)
		}

		fmt.Printf("Produced msg %v key: %v\n", i, msg.Key)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("All msg produced")
}
