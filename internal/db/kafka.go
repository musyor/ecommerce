package db

import (
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	Writer *kafka.Writer
}

//func NewProducer(cfg *config.Config, topic string) *Producer {
//	addr := kafka.TCP(fmt.Sprintf("%s:%s", cfg.KafkaHost, cfg.KafkaPort))
//
//	writer := &kafka.Writer{
//		Addr:         addr,
//		Topic:        topic,
//		Balancer:     ,
//		Compression:  ,
//		RequiredAcks: ,
//	}
//	return &Producer{Writer: writer}
//}
