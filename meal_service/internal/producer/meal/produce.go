package mealproducer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/JoePeach762/PP_project/meal_service/internal/models"
	"github.com/segmentio/kafka-go"
)

func (p *MealKafkaProducer) PublishMealConsumed(ctx context.Context, meal *models.MealInfo) error {
	data, err := json.Marshal(meal)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Value: data,
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return p.writer.WriteMessages(ctx, msg)
}
