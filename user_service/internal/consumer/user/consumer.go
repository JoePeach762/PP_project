package userconsumer

import (
	"context"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
)

type processor interface {
	AddMealToUser(ctx context.Context, mealInfo *models.MealInfo) error
}

type Consumer struct {
	processor processor
	kafka     []string
	topic     string
}

func NewUserConsumer(processor processor, kafka []string, topic string) *Consumer {
	return &Consumer{
		processor: processor,
		kafka:     kafka,
		topic:     topic,
	}
}
