package models

// SmsMessage initializes the SMS message object
type SmsMessage struct {
	Timestamp string
	Sender    string
	Receiver  string
	Message   string
}

// GetSmsMessage ...
func GetSmsMessage(smsMessage []string) SmsMessage {
	smsMessageObject := SmsMessage{
		Timestamp: smsMessage[0],
		Sender:    smsMessage[1],
		Receiver:  smsMessage[2],
		Message:   smsMessage[3],
	}

	return smsMessageObject
}
