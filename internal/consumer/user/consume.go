package userconsumer

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/JoePeach762/PP_project/internal/models"
	"github.com/segmentio/kafka-go"
)

func (c *consumer) Consume(ctx context.Context) {
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
		var userInfo *models.UserInfo
		err = json.Unmarshal(msg.Value, &userInfo)
		if err != nil {
			slog.Error("parce", "error", err)
			continue
		}
		err = c.processor.Handle(ctx, userInfo)
		if err != nil {
			slog.Error("Handle", "error", err)
		}
	}

}
