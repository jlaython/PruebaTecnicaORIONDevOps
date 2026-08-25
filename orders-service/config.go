package main

import "os"

type Config struct {
	ServerPort   string
	RabbitURL    string
	QueueName    string
	RedisAddr    string
	RedisPass    string
	RedisListKey string
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func LoadConfig() Config {
	return Config{
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		RabbitURL:    getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		QueueName:    getEnv("RABBITMQ_QUEUE", "events_queue"),
		RedisAddr:    getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPass:    getEnv("REDIS_PASSWORD", ""),
		RedisListKey: getEnv("REDIS_LIST_KEY", "events_list"),
	}
}
