package bootstrap

import (
	"log"

	"github.com/JoePeach762/PP_project/config"
	mealproducer "github.com/JoePeach762/PP_project/internal/producer/meal"
	"github.com/segmentio/kafka-go"
)

func InitKafkaProducer(cfg *config.Config) *mealproducer.MealKafkaProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(cfg.Kafka.Brokers...),
		Topic:    cfg.Kafka.Topics.MealConsumed,
		Balancer: &kafka.LeastBytes{},
		ErrorLogger: kafka.LoggerFunc(func(msg string, args ...interface{}) {
			log.Printf("Kafka error: "+msg, args...)
		}),
	}
	return mealproducer.NewKafkaProducer(writer)
}
