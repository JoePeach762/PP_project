package userproducer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

func (p *UserKafkaProducer) PublishUsersDeleted(ctx context.Context, ids []uint64) error {
	data, err := json.Marshal(ids)
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
