package models

// MessageEnriched initializes the enriched message object
type MessageEnriched struct {
	Timestamp int64
	Sender    string
	Receiver  string
	Message   string
	HamOrSpam bool
	Accuracy  string
}
