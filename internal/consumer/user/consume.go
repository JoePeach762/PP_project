package userconsumer

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/JoePeach762/PP_project/internal/models"
	"github.com/segmentio/kafka-go"
)

type MealConsumedEvent struct {
	UserID       uint64    `json:"user_id"`
	Name         string    `json:"name"`
	WeightGrams  float32   `json:"weight_grams"`
	Calories100g float32   `json:"calories_100g"`
	Proteins100g float32   `json:"proteins_100g"`
	Fats100g     float32   `json:"fats_100g"`
	Carbs100g    float32   `json:"carbs_100g"`
	Date         time.Time `json:"date"`
}

func (c *Consumer) Consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           c.kafka,
		GroupID:           "UserService_group",
		Topic:             c.topic,
		HeartbeatInterval: 3 * time.Second,
		SessionTimeout:    30 * time.Second,
	})
	defer r.Close()

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			slog.Error("consumer.consume error", "error", err.Error())
		}
		var event MealConsumedEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			slog.Error("parse event", "error", err)
			continue
		}

		mealInfo := &models.MealInfo{
			Name:         event.Name,
			WeightGrams:  event.WeightGrams,
			Calories100g: event.Calories100g,
			Proteins100g: event.Proteins100g,
			Fats100g:     event.Fats100g,
			Carbs100g:    event.Carbs100g,
			Date:         event.Date,
		}

		if err := c.processor.AddMealToUser(ctx, event.UserID, mealInfo); err != nil {
			slog.Error("AddMealToUser", "error", err)
			// TODO: повторный вызов или отправка в DLQ
		}
	}

}
