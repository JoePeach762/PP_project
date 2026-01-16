package bootstrap

import (
	"github.com/JoePeach762/PP_project/user_service/config"
	userconsumer "github.com/JoePeach762/PP_project/user_service/internal/consumer/user"
	userprocessor "github.com/JoePeach762/PP_project/user_service/internal/processors/user"
)

func InitUserConsumer(cfg *config.Config, processor *userprocessor.Processor) *userconsumer.Consumer {
	return userconsumer.NewUserConsumer(processor, cfg.Kafka.Brokers, cfg.Kafka.MealConsumedTopicName)
}
