package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
)

func main() {
	brokerAddr := "localhost:9092"
	topic := "test-topic"
	groupID := "test-consumer-group"

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerAddr},
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 1,
		MaxBytes: 10e6,
	})

	defer reader.Close()

	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("Shutting down consumer")
		cancel()
	}()

	fmt.Printf("Consumer listening on topic '%s'... (Ctrl+C to stop)\n\n", topic)

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				break
			}
			log.Printf("Error reading msg: %v", err)
			continue
		}

		fmt.Printf("Consumed: partition=%d, offset:%d, key:%s, value:%s\n",
			msg.Partition, msg.Offset, msg.Key, msg.Value)
	}
	fmt.Println("Consumer stopped.")
}
