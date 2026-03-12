package userconsumer

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log/slog"
	"time"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/segmentio/kafka-go"
)

const consumerGroupID = "user-service-consumer"

func (c *Consumer) Consume(ctx context.Context) {
	slog.Info("Starting Kafka consumer", "topic", c.topic, "group_id", consumerGroupID)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           c.kafka,
		GroupID:           consumerGroupID,
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
		msg, err := r.FetchMessage(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				slog.Info("Kafka consumer stopped")
				return
			}

			slog.Error("Failed to fetch Kafka message", "error", err)

			time.Sleep(time.Second)
			continue
		}

		var event models.MealInfo
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			slog.Error("Failed to parse Kafka message",
				"error", err,
				"message_offset", msg.Offset,
				"message_partition", msg.Partition)

			if err := commitMessage(ctx, r, msg); err != nil {
				slog.Error("Failed to commit invalid Kafka message",
					"error", err,
					"message_offset", msg.Offset,
					"message_partition", msg.Partition)
			}
			continue
		}

		if event.UserId == 0 {
			slog.Warn("Received meal event with missing UserID",
				"message_offset", msg.Offset,
				"meal_name", event.Name)

			if err := commitMessage(ctx, r, msg); err != nil {
				slog.Error("Failed to commit Kafka message with missing UserID",
					"error", err,
					"message_offset", msg.Offset,
					"message_partition", msg.Partition)
			}
			continue
		}

		if event.EventID == "" {
			event.EventID = deriveLegacyEventID(&event)
		}

		if err := c.processor.AddMealToUser(ctx, &event); err != nil {
			if isNonRetryableProcessingError(err) {
				slog.Warn("Skipping non-retryable meal event",
					"error", err,
					"user_id", event.UserId,
					"meal_name", event.Name,
					"message_offset", msg.Offset,
					"message_partition", msg.Partition)

				if err := commitMessage(ctx, r, msg); err != nil {
					slog.Error("Failed to commit non-retryable Kafka message",
						"error", err,
						"message_offset", msg.Offset,
						"message_partition", msg.Partition)
				}
				continue
			}

			slog.Error("Failed to process meal event",
				"error", err,
				"user_id", event.UserId,
				"meal_name", event.Name,
				"message_offset", msg.Offset)
			continue
		}

		if err := commitMessage(ctx, r, msg); err != nil {
			slog.Error("Failed to commit processed Kafka message",
				"error", err,
				"user_id", event.UserId,
				"meal_name", event.Name,
				"message_offset", msg.Offset,
				"message_partition", msg.Partition)
			continue
		}

		slog.Debug("Successfully processed meal event",
			"user_id", event.UserId,
			"meal_name", event.Name,
			"calories", event.Calories100g*event.WeightGrams/100)
	}
}

func commitMessage(ctx context.Context, reader *kafka.Reader, msg kafka.Message) error {
	return reader.CommitMessages(ctx, msg)
}

func isNonRetryableProcessingError(err error) bool {
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return false
	}

	return pgErr.Code == "23503"
}

func deriveLegacyEventID(event *models.MealInfo) string {
	payload, err := json.Marshal(event)
	if err != nil {
		return ""
	}

	sum := sha256.Sum256(payload)
	return hex.EncodeToString(sum[:])
}
