package userconsumer

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

type processor interface {
	AddMealToUser(ctx context.Context, id uint64, mealInfo *models.MealInfo) error
}

type consumer struct {
	processor processor
	kafka     []string
	topic     string
}

func NewUserConsumer(processor processor, kafka []string, topic string) *consumer {
	return &consumer{
		processor: processor,
		kafka:     kafka,
		topic:     topic,
	}
}
