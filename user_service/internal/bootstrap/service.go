package bootstrap

import (
	"github.com/JoePeach762/PP_project/user_service/config"
	userproducer "github.com/JoePeach762/PP_project/user_service/internal/producer/user"
	"github.com/JoePeach762/PP_project/user_service/internal/services/user"
	userstorage "github.com/JoePeach762/PP_project/user_service/internal/storage/pgstorage/userstorage"
)

func InitUserService(
	storage *userstorage.PGstorage,
	producer *userproducer.UserKafkaProducer,
	cfg *config.Config,
) *user.Service {
	return user.NewUserService(
		storage,
		producer,
		uint32(cfg.UserServiceSettings.MinNameLen),
		uint32(cfg.UserServiceSettings.MaxNameLen),
		uint32(cfg.UserServiceSettings.MinWeight),
		uint32(cfg.UserServiceSettings.MaxWeight),
	)
}
