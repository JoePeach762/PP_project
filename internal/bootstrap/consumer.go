package bootstrap

import (
	"github.com/JoePeach762/PP_project/config"
	userconsumer "github.com/JoePeach762/PP_project/internal/consumer/user"
	userprocessor "github.com/JoePeach762/PP_project/internal/services/processors/user"
)

func InitUserConsumer(cfg *config.Config, processor *userprocessor.Processor) *userconsumer.Consumer {
	brokers := cfg.Kafka.Brokers

	return userconsumer.NewUserConsumer(processor, brokers, cfg.Kafka.Topics.MealConsumed)
}
