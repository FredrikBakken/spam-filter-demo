package kafka

import (
	"encoding/json"
	"fmt"

	"telenor.com/spam-filter-demo/sms-event-integration/models"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"telenor.com/spam-filter-demo/sms-event-integration/config"
)

// ProduceMessage sends the SMS message to a Kafka topic
func ProduceMessage(cfg *config.Config, producer *kafka.Producer, message models.Message, key string) {
	deliveryChan := make(chan kafka.Event)

	value, _ := json.Marshal(&message)
	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &cfg.Kafka.Topic,
			Partition: kafka.PartitionAny},
		Key:   []byte(key),
		Value: []byte(value),
	}, deliveryChan)

	if err != nil {
		fmt.Println("Error creating producer message...")
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
}
