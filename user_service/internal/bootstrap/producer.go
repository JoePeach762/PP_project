package bootstrap

import (
	"github.com/JoePeach762/PP_project/user_service/config"
	userproducer "github.com/JoePeach762/PP_project/user_service/internal/producer/user"
	"github.com/segmentio/kafka-go"
)

func InitKafkaProducer(cfg *config.Config) *userproducer.UserKafkaProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(cfg.Kafka.Brokers...),
		Topic:    cfg.Kafka.UserDeletedTopicName,
		Balancer: &kafka.LeastBytes{},
	}
	return userproducer.NewKafkaProducer(writer)
}
