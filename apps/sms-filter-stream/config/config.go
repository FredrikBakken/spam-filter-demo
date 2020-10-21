package config

import (
	"os"
	"strings"
)

// ServerConfig defines the port and event-route
type ServerConfig struct {
	Route string
	Port  string
}

// KafkaConfig defines the Kafka-related variables
type KafkaConfig struct {
	Broker      []string
	TopicNewSMS string
	TopicHam    string
	TopicSpam   string
}

// ProducerConfig defines the configuration of the Kafka producer
type ProducerConfig struct {
	ClientID        string
	Acks            string
	Retries         string
	CompressionType string
}

// Config is the collection of sub-configs
type Config struct {
	Server   ServerConfig
	Kafka    KafkaConfig
	Producer ProducerConfig
}

// New builds the application configuration
func New() *Config {
	return &Config{
		Server: ServerConfig{
			Route: getEnv("API_URL", "/event"),
			Port:  getEnv("PORT", "9090"),
		},
		Kafka: KafkaConfig{
			Broker:      getEnvAsSlice("BOOTSTRAP_SERVERS", []string{"localhost:9092"}),
			TopicNewSMS: getEnv("TOPIC_NEW_SMS", "new-sms-json-v1"),
			TopicHam:    getEnv("TOPIC_HAM_SMS", "safe-sms-json-v1"),
			TopicSpam:   getEnv("TOPIC_SPAM_SMS", "spam-sms-json-v1"),
		},
		Producer: ProducerConfig{
			ClientID:        getEnv("CLIENT_ID", "sms-filter-stream"),
			Acks:            getEnv("ACKS", "all"),
			Retries:         getEnv("RETRIES", "10"),
			CompressionType: getEnv("COMPRESSION_TYPE", "lz4"),
		},
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsSlice(key string, fallback []string) []string {
	valueString := getEnv(key, "")
	if valueString != "" {
		slice := strings.Split(valueString, ",")
		return slice
	}
	return fallback
}
