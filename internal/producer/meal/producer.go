package mealproducer

import "github.com/segmentio/kafka-go"

type producer struct {
	brokers []string
	writer  *kafka.Writer
}
