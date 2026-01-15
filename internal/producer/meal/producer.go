package mealproducer

import (
	"github.com/segmentio/kafka-go"
)

type MealKafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(writer *kafka.Writer) *MealKafkaProducer {
	return &MealKafkaProducer{writer: writer}
}
