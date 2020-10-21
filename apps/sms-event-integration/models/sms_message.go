package models

// SmsMessage initializes the SMS message object
type SmsMessage struct {
	Timestamp string
	Sender    string
	Receiver  string
	Message   string
}
