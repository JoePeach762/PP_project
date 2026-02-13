package bootstrap

import (
	"github.com/JoePeach762/PP_project/meal_service/config"
	mealconsumer "github.com/JoePeach762/PP_project/meal_service/internal/consumer/meal"
	mealprocessor "github.com/JoePeach762/PP_project/meal_service/internal/processors/meal"
)

func InitMealConsumer(cfg *config.Config, processor *mealprocessor.Processor) *mealconsumer.Consumer {
	return mealconsumer.NewMealConsumer(processor, cfg.Kafka.Brokers, cfg.Kafka.UserDeletedTopicName)
}
