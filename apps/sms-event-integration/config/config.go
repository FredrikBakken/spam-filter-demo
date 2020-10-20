package config

import "os"

// ServerConfig defines the port and event-route
type ServerConfig struct {
	Route string
	Port  string
}

// KafkaConfig defines the Kafka-related variables
type KafkaConfig struct {
	Broker     string
	Schema     string
	Topic      string
	SchemaType string
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
			Broker:     getEnv("BOOTSTRAP_SERVERS", "localhost:9092"),
			Schema:     getEnv("SCHEMA_REGISTRY", "http://localhost:8081"),
			Topic:      getEnv("KAFKA_TOPIC", "new-sms-json-v1"),
			SchemaType: getEnv("SCHEMA_TYPE", "JSON"),
		},
		Producer: ProducerConfig{
			ClientID:        getEnv("CLIENT_ID", "sms-filter-integration"),
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
