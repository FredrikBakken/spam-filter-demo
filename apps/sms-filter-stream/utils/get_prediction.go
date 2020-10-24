package utils

import (
	"bytes"
	"encoding/json"
	"net/http"

	"telenor.com/spam-filter-demo/sms-filter-stream/models"
)

// GetPrediction ...
func GetPrediction(message models.Message) models.MessageEnriched {
	url := "http://localhost:5000/sms"

	jsonStr := []byte(`{"message": "` + message.Message + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var prediction models.Prediction
	json.NewDecoder(resp.Body).Decode(&prediction)

	var messageEnriched = models.MessageEnriched{
		Timestamp: message.Timestamp,
		Sender:    message.Sender,
		Receiver:  message.Receiver,
		Message:   message.Message,
		HamOrSpam: prediction.Spam,
		Accuracy:  prediction.Confidence,
	}

	return messageEnriched
}
