package mealconsumer

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"time"

	"github.com/segmentio/kafka-go"
)

func (c *Consumer) Consume(ctx context.Context) {
	slog.Info("Starting Kafka consumer",
		"topic", c.topic,
		"group_id", "meal-service-consumer",
	)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           c.kafka,
		GroupID:           "meal-service-consumer",
		Topic:             c.topic,
		MinBytes:          1,
		MaxBytes:          10e6,
		HeartbeatInterval: 3 * time.Second,
		SessionTimeout:    30 * time.Second,
		RebalanceTimeout:  10 * time.Second,
	})
	defer r.Close()

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
			continue
		}

		if err := c.processor.CascadeDeletion(ctx, ids); err != nil {
			slog.Error("Processing failed",
				"error", err,
				"offset", msg.Offset,
			)
			continue
		}

		if err := r.CommitMessages(ctx, msg); err != nil {
			slog.Error("Commit failed",
				"error", err,
				"offset", msg.Offset,
			)
			continue
		}

		slog.Debug("User meals deleted", "ids", ids)
	}
}
