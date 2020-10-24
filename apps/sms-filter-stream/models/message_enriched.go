package models

// MessageEnriched initializes the enriched message object
type MessageEnriched struct {
	Timestamp string
	Sender    string
	Receiver  string
	Message   string
	HamOrSpam bool
	Accuracy  string
}
