package kafka

import (
	"log"

	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
	"telenor.com/spam-filter-demo/sms-filter-stream/config"
)

// CreateProcessor ...
func CreateProcessor(cfg *config.Config, callback func(goka.Context, interface{})) (*goka.Processor, error) {
	gFilter := goka.DefineGroup(goka.Group(cfg.Kafka.TopicHam),
		goka.Input(goka.Stream(cfg.Kafka.TopicNewSMS), new(codec.String), callback),
		goka.Output(goka.Stream(cfg.Kafka.TopicHam), new(codec.Bytes)),
		goka.Output(goka.Stream(cfg.Kafka.TopicSpam), new(codec.Bytes)),
		goka.Persist(new(codec.Bytes)),
	)

	// Create new m2m-filter process
	processor, err := goka.NewProcessor(cfg.Kafka.Broker, gFilter)
	if err != nil {
		log.Fatalf("Error creating the m2m-filter processor: %v", err)
	}

	return processor, err
}
