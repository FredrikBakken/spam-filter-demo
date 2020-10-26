package firestore

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"telenor.com/spam-filter-demo/sms-event-integration/models"
	"telenor.com/spam-filter-demo/sms-event-integration/utils"
)

var (
	forwardedMessages = []string{}
)

// GetMessages ...
func GetMessages() []models.Message {
	var newMessages = []models.Message{}

	// Setup Firestore connection
	opt := option.WithCredentialsFile("config/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// Read from the Cloud Firestore
	docs := client.Collection("messages").Documents(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// Loop through the found Documents
	for {
		doc, err := docs.Next()

		// No more Documents to loop through
		if err == iterator.Done {
			break
		}

		// New message
		if !utils.ExistsInSlice(forwardedMessages, doc.Ref.ID) {
			log.Println(doc.Data())

			var newMessage = models.Message{
				Timestamp: doc.Data()["timestamp"].(int64),
				Sender:    doc.Data()["username"].(string),
				Receiver:  doc.Ref.ID,
				Message:   doc.Data()["message"].(string),
			}

			// Append new messages
			newMessages = append(newMessages, newMessage)

			// Append messages already read
			forwardedMessages = append(forwardedMessages, doc.Ref.ID)
		}
	}

	// Close the client connection
	defer client.Close()

	return newMessages
}
