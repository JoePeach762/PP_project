package userconsumer

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
	"github.com/segmentio/kafka-go"
)

func (c *Consumer) Consume(ctx context.Context) {
	slog.Info("Starting Kafka consumer", "topic", c.topic, "group_id", "user-service-consumer")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           c.kafka,
		GroupID:           "user-service-consumer",
		Topic:             c.topic,
		MinBytes:          1,
		MaxBytes:          10e6,
		HeartbeatInterval: 3 * time.Second,
		SessionTimeout:    30 * time.Second,
		RebalanceTimeout:  10 * time.Second,
	})
	defer func() {
		if err := r.Close(); err != nil {
			slog.Error("Failed to close Kafka reader", "error", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			slog.Info("Kafka consumer stopped")
			return
		default:
		}

		msg, err := r.ReadMessage(ctx)

		if err != nil {
			if err == context.Canceled || err == context.DeadlineExceeded {
				continue
			}

			slog.Error("Failed to read Kafka message", "error", err)

			time.Sleep(1 * time.Second)
			continue
		}

		var event models.MealInfo
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			slog.Error("Failed to parse Kafka message",
				"error", err,
				"message_offset", msg.Offset,
				"message_partition", msg.Partition)
			continue
		}

		if event.UserId == 0 {
			slog.Warn("Received meal event with missing UserID",
				"message_offset", msg.Offset,
				"meal_name", event.Name)
			continue
		}

		if err := c.processor.AddMealToUser(ctx, &event); err != nil {
			slog.Error("Failed to process meal event",
				"error", err,
				"user_id", event.UserId,
				"meal_name", event.Name,
				"message_offset", msg.Offset)
			continue
		}

		slog.Debug("Successfully processed meal event",
			"user_id", event.UserId,
			"meal_name", event.Name,
			"calories", event.Calories100g*event.WeightGrams/100)
	}
}
