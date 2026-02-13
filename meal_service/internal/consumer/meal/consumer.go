package mealconsumer

import (
	"context"
)

type processor interface {
	CascadeDeletion(ctx context.Context, ids []uint64) error
}

type Consumer struct {
	processor processor
	kafka     []string
	topic     string
}

func NewMealConsumer(processor processor, kafka []string, topic string) *Consumer {
	return &Consumer{
		processor: processor,
		kafka:     kafka,
		topic:     topic,
	}
}
