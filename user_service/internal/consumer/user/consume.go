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
	slog.Info("Запуск консьюмера Kafka", "topic", c.topic, "group_id", consumerGroupID)

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
			slog.Error("Не удалось закрыть ридер Kafka", "error", err)
		}
	}()

	for {
		msg, err := r.FetchMessage(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				slog.Info("Консьюмер Kafka остановлен")
				return
			}

			slog.Error("Не удалось получить сообщение из Kafka", "error", err)

			time.Sleep(time.Second)
			continue
		}

		var event models.MealInfo
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			slog.Error("Не удалось разобрать сообщение Kafka",
				"error", err,
				"message_offset", msg.Offset,
				"message_partition", msg.Partition)

			if err := commitMessage(ctx, r, msg); err != nil {
				slog.Error("Не удалось закоммитить некорректное сообщение Kafka",
					"error", err,
					"message_offset", msg.Offset,
					"message_partition", msg.Partition)
			}
			continue
		}

		if event.UserId == 0 {
			slog.Warn("Получено событие о приеме пищи без UserID",
				"message_offset", msg.Offset,
				"meal_name", event.Name)

			if err := commitMessage(ctx, r, msg); err != nil {
				slog.Error("Не удалось закоммитить сообщение Kafka без UserID",
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
				slog.Warn("Пропуск события о приеме пищи с неповторяемой ошибкой",
					"error", err,
					"user_id", event.UserId,
					"meal_name", event.Name,
					"message_offset", msg.Offset,
					"message_partition", msg.Partition)

				if err := commitMessage(ctx, r, msg); err != nil {
					slog.Error("Не удалось закоммитить сообщение Kafka с неповторяемой ошибкой",
						"error", err,
						"message_offset", msg.Offset,
						"message_partition", msg.Partition)
				}
				continue
			}

			slog.Error("Не удалось обработать событие о приеме пищи",
				"error", err,
				"user_id", event.UserId,
				"meal_name", event.Name,
				"message_offset", msg.Offset)
			continue
		}

		if err := commitMessage(ctx, r, msg); err != nil {
			slog.Error("Не удалось закоммитить обработанное сообщение Kafka",
				"error", err,
				"user_id", event.UserId,
				"meal_name", event.Name,
				"message_offset", msg.Offset,
				"message_partition", msg.Partition)
			continue
		}

		slog.Debug("Событие о приеме пищи успешно обработано",
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
