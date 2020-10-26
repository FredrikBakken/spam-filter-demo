package firestore

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// UpdateMessages ...
func UpdateMessages(docID string, message map[string]interface{}) {
	// Setup Firestore connection
	opt := option.WithCredentialsFile("config/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// Write message update to Cloud Firestore
	result, err := client.Collection("messages").Doc(docID).Set(context.Background(), message)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)
}
