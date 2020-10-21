package kafka

import (
	"encoding/json"

	"telenor.com/spam-filter-demo/sms-filter-stream/utils"

	"github.com/lovoo/goka"
	"telenor.com/spam-filter-demo/sms-filter-stream/config"
	"telenor.com/spam-filter-demo/sms-filter-stream/models"
)

// CallbackFilter ...
func CallbackFilter(cfg *config.Config) func(goka.Context, interface{}) {
	return func(ctx goka.Context, msg interface{}) {
		var (
			sms         models.Sms
			smsEnriched models.SmsEnriched
		)

		// Unmarshal the data from the incoming sms message
		err := json.Unmarshal([]byte(msg.(string)), &sms)
		if err != nil {
			panic(err)
		}

		// Get the spam-predicion
		smsEnriched = utils.GetPrediction(sms)
		value, _ := json.Marshal(&smsEnriched)

		// True: Message is spam
		// False: Message is not spam
		if smsEnriched.HamOrSpam == true {
			ctx.Emit(goka.Stream(cfg.Kafka.TopicSpam), smsEnriched.Sender, value)
		} else {
			ctx.Emit(goka.Stream(cfg.Kafka.TopicHam), smsEnriched.Sender, value)
		}
	}
}
