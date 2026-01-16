package bootstrap

import (
	"github.com/JoePeach762/PP_project/user_service/config"
	userconsumer "github.com/JoePeach762/PP_project/user_service/internal/consumer/user"
	userprocessor "github.com/JoePeach762/PP_project/user_service/internal/processors/user"
)

func InitUserConsumer(cfg *config.Config, processor *userprocessor.Processor) *userconsumer.Consumer {
	brokers := cfg.Kafka.Brokers

	return userconsumer.NewUserConsumer(processor, brokers, cfg.Kafka.Topics.MealConsumed)
}
