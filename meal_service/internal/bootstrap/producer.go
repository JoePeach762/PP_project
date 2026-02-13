package bootstrap

import (
	"github.com/JoePeach762/PP_project/meal_service/config"
	mealproducer "github.com/JoePeach762/PP_project/meal_service/internal/producer/meal"
	"github.com/segmentio/kafka-go"
)

func InitKafkaProducer(cfg *config.Config) *mealproducer.MealKafkaProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(cfg.Kafka.Brokers...),
		Topic:    cfg.Kafka.MealConsumedTopicName,
		Balancer: &kafka.LeastBytes{},
	}
	return mealproducer.NewKafkaProducer(writer)
}
