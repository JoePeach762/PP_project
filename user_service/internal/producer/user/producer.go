package userproducer

import (
	"github.com/segmentio/kafka-go"
)

type UserKafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(writer *kafka.Writer) *UserKafkaProducer {
	return &UserKafkaProducer{writer: writer}
}
