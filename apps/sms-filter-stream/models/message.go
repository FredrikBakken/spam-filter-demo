package models

// Message initializes the message object
type Message struct {
	Timestamp int64
	Sender    string
	Receiver  string
	Message   string
}
