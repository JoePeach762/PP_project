package userproducer

import (
	"context"
	"encoding/json"

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

	return p.PublishMessage(ctx, msg.Value)
}
