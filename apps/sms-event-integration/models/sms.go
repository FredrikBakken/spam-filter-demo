package models

// Sms initializes the SMS message object
type Sms struct {
	Timestamp string
	Sender    string
	Receiver  string
	Message   string
}
