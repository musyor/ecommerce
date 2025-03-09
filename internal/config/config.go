package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	MySQLHost     string
	MySQLPort     string
	MySQLUser     string
	MySQLPassword string
	MySQLDatabase string

	RedisPort     string
	RedisHost     string
	RedisPassword string

	KafkaHost string
	KafkaPort string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	cfg.MySQLHost = getEnv("MYSQL_HOST", "localhost")
	cfg.MySQLPort = getEnv("MYSQL_PORT", "3306")
	cfg.MySQLUser = getEnv("MYSQL_USER", "root")
	cfg.MySQLPassword = getEnv("MYSQL_PASSWORD", "root")
	cfg.MySQLDatabase = getEnv("MYSQL_DATABASE", "ecommerce")

	// Redis 配置
	cfg.RedisHost = getEnv("REDIS_HOST", "localhost")
	cfg.RedisPort = getEnv("REDIS_PORT", "6379")
	cfg.RedisPassword = getEnv("REDIS_PASSWORD", "")

	// Kafka 配置
	cfg.KafkaHost = getEnv("KAFKA_HOST", "localhost")
	cfg.KafkaPort = getEnv("KAFKA_PORT", "9092")

	return cfg, nil
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
