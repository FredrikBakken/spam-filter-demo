package handlers

import (
	"encoding/json"
	"net/http"

	"telenor.com/spam-filter-demo/sms-event-integration/config"
	"telenor.com/spam-filter-demo/sms-event-integration/models"
	"telenor.com/spam-filter-demo/sms-event-integration/services/kafka"
)

// NewSmsEvent forwards the incoming request by concurrency
func NewSmsEvent(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var sms models.Message
	err := json.NewDecoder(req.Body).Decode(&sms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go ConcurrencyHandler(sms)
}

// ConcurrencyHandler handles the process of passing data to Kafka in parallel
func ConcurrencyHandler(sms models.Message) {
	var (
		cfg      = config.New()
		producer = kafka.CreateKafkaProducer(cfg)
	)

	kafka.ProduceMessage(cfg, producer, sms, "")

	producer.Close()
}
