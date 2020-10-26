package kafka

import (
	"encoding/json"
	"strconv"

	"telenor.com/spam-filter-demo/sms-filter-stream/services/firestore"
	"telenor.com/spam-filter-demo/sms-filter-stream/utils"

	"github.com/lovoo/goka"
	"telenor.com/spam-filter-demo/sms-filter-stream/config"
	"telenor.com/spam-filter-demo/sms-filter-stream/models"
)

// CallbackFilter ...
func CallbackFilter(cfg *config.Config) func(goka.Context, interface{}) {
	return func(ctx goka.Context, msg interface{}) {
		var (
			message         models.Message
			messageEnriched models.MessageEnriched
		)

		// Unmarshal the data from the incoming sms message
		err := json.Unmarshal([]byte(msg.(string)), &message)
		if err != nil {
			panic(err)
		}

		// Get the spam-predicion
		messageEnriched = utils.GetPrediction(message)
		value, _ := json.Marshal(&messageEnriched)

		// True: Message is spam
		// False: Message is not spam
		if messageEnriched.HamOrSpam == true {
			ctx.Emit(goka.Stream(cfg.Kafka.TopicSpam), messageEnriched.Sender, value)
		} else {
			ctx.Emit(goka.Stream(cfg.Kafka.TopicHam), messageEnriched.Sender, value)
		}

		// Send to Firestore
		if ctx.Key() != "" {
			enrichedMsg := make(map[string]interface{})
			enrichedMsg["timestamp"] = messageEnriched.Timestamp
			enrichedMsg["username"] = messageEnriched.Sender
			enrichedMsg["message"] = messageEnriched.Message
			enrichedMsg["ham-or-spam"] = strconv.FormatBool(messageEnriched.HamOrSpam)
			enrichedMsg["accuracy"] = messageEnriched.Accuracy
			firestore.UpdateMessages(ctx.Key(), enrichedMsg)
		}
	}
}
