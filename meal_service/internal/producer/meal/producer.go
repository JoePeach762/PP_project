package mealproducer

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type MealKafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(writer *kafka.Writer) *MealKafkaProducer {
	return &MealKafkaProducer{writer: writer}
}

func (p *MealKafkaProducer) PublishMessage(ctx context.Context, payload []byte) error {
	msg := kafka.Message{
		Value: payload,
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return p.writer.WriteMessages(ctx, msg)
}

func (p *MealKafkaProducer) Close() error {
	return p.writer.Close()
}
