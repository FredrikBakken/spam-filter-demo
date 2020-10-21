package models

// SmsEnriched initializes the enriched SMS message object
type SmsEnriched struct {
	Timestamp string
	Sender    string
	Receiver  string
	Message   string
	HamOrSpam bool
	Accuracy  string
}
