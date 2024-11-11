package config

type Config struct {
	KafkaBootstrapServers string
	OrdersTopic           string
}

func LoadConfig() Config {
	return Config{
		KafkaBootstrapServers: "127.0.0.1:49388",
		OrdersTopic:           "orders",
	}
}
