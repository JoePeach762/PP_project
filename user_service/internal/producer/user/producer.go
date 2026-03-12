package userproducer

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type UserKafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(writer *kafka.Writer) *UserKafkaProducer {
	return &UserKafkaProducer{writer: writer}
}

func (p *UserKafkaProducer) PublishMessage(ctx context.Context, payload []byte) error {
	msg := kafka.Message{
		Value: payload,
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return p.writer.WriteMessages(ctx, msg)
}

func (p *UserKafkaProducer) Close() error {
	return p.writer.Close()
}
