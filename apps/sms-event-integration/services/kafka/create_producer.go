package kafka

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"telenor.com/spam-filter-demo/sms-event-integration/config"
)

// CreateKafkaProducer initializes and returns a Kafka Producer
func CreateKafkaProducer(cfg *config.Config) *kafka.Producer {
	config := &kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.Broker,
		"acks":              cfg.Producer.Acks,
		"retries":           cfg.Producer.Retries,
		"compression.type":  cfg.Producer.CompressionType,
		"client.id":         cfg.Producer.ClientID,
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Println(fmt.Sprintf("Error creating the Kafka Producer: %s", err))
	}

	return producer
}
