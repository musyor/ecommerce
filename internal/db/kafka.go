package db

import (
	"ecommerce/internal/config"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	Writer *kafka.Writer
}

func NewProducer(cfg *config.Config, topic string) *Producer {
	addr := kafka.TCP(fmt.Sprintf("%s:%s", cfg.KafkaHost, cfg.KafkaPort))

	writer := &kafka.Writer{
		Addr:         addr,
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{}, // 指定分区的balancer模式为最小字节分布
		RequiredAcks: kafka.RequireAll,    // ack模式
		Async:        true,                // 异步
	}
	return &Producer{Writer: writer}
}
