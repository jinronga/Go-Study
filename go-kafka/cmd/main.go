package main

import (
	"github.com/IBM/sarama"
	"log"
)

type (
	testMQ struct {
	}
)

func main() {
	t := &testMQ{}
	t.testKafka()
}

func (t *testMQ) testKafka() {
	// Kafka broker addresses
	brokers := []string{"localhost:9092"}
	topic := "quickstart-events"

	// Sarama configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll

	// Create producer
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	// Send message
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("0"),
	}
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}
