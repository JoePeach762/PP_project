package mealconsumer

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"time"

	"github.com/segmentio/kafka-go"
)

const consumerGroupID = "meal-service-consumer"

func (c *Consumer) Consume(ctx context.Context) {
	slog.Info("Starting Kafka consumer",
		"topic", c.topic,
		"group_id", consumerGroupID,
	)

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

		var ids []uint64
		if err := json.Unmarshal(msg.Value, &ids); err != nil {
			slog.Error("Invalid Kafka payload",
				"error", err,
				"offset", msg.Offset,
				"partition", msg.Partition,
			)

			if err := commitMessage(ctx, r, msg); err != nil {
				slog.Error("Failed to commit invalid Kafka message",
					"error", err,
					"offset", msg.Offset,
					"partition", msg.Partition,
				)
			}
			continue
		}

		if err := c.processor.CascadeDeletion(ctx, ids); err != nil {
			slog.Error("Processing failed",
				"error", err,
				"offset", msg.Offset,
			)
			continue
		}

		if err := commitMessage(ctx, r, msg); err != nil {
			slog.Error("Commit failed",
				"error", err,
				"offset", msg.Offset,
				"partition", msg.Partition,
			)
			continue
		}

		slog.Debug("User meals deleted", "ids", ids)
	}
}

func commitMessage(ctx context.Context, reader *kafka.Reader, msg kafka.Message) error {
	return reader.CommitMessages(ctx, msg)
}
