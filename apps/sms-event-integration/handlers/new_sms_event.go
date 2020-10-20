package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"telenor.com/spam-filter-demo/sms-event-integration/config"
	"telenor.com/spam-filter-demo/sms-event-integration/models"
	"telenor.com/spam-filter-demo/sms-event-integration/services/kafka"
)

// NewSmsEvent forwards the incoming request by concurrency
func NewSmsEvent(_ http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	go ConcurrencyHandler(body)
}

// ConcurrencyHandler handles the process of passing data to Kafka in parallel
func ConcurrencyHandler(body []byte) {
	var (
		cfg      = config.New()
		messages = bytes.Split(body, []byte("\n"))

		// Initialize the Kafka Producer
		producer = kafka.CreateKafkaProducer(cfg)
	)

	for _, message := range messages {
		var smsMessage = strings.Split(string(message), ",")
		fmt.Println(smsMessage)

		newMessage := models.GetSmsMessage(smsMessage)
		kafka.ProduceMessage(producer, newMessage)
	}

	producer.Close()
}
