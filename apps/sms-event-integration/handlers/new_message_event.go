package handlers

import (
	"time"

	"telenor.com/spam-filter-demo/sms-event-integration/config"
	"telenor.com/spam-filter-demo/sms-event-integration/services/firestore"
	"telenor.com/spam-filter-demo/sms-event-integration/services/kafka"
)

// NewMessageEvent ...
func NewMessageEvent(cfg *config.Config) {
	for {
		// fmt.Println("Running right now!")

		// Get the newest messages
		var newMessages = firestore.GetMessages()

		// Configure producer
		var producer = kafka.CreateKafkaProducer(cfg)

		for _, newMessage := range newMessages {
			kafka.ProduceMessage(cfg, producer, newMessage, newMessage.Receiver)
		}

		producer.Close()

		time.Sleep(10000 * time.Millisecond)
	}
}
