package db

import (
	"ecommerce/internal/config"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Client *redis.Client // 公开的字段
}

var RedisInstance *Redis

func NewRedis(cfg *config.Config) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       0,
	})
	return &Redis{Client: client}
}
