package config

import "github.com/caarlos0/env/v6"

type Config struct {
	MySQLHost     string
	MySQLPort     string
	MySQLUser     string
	MySQLPassword string
	MySQLDatabase string

	RedisPort     int
	RedisHost     string
	RedisPassword string

	KafkaHost string
	KafkaPort int
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
