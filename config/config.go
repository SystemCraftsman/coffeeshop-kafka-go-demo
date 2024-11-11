
package config

type Config struct {
    KafkaBootstrapServers string
    CoffeeshopTopic       string
}

func LoadConfig() Config {
    return Config{
        KafkaBootstrapServers: "localhost:9092",
        CoffeeshopTopic:       "orders",
    }
}
