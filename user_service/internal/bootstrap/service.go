package bootstrap

import (
	"github.com/JoePeach762/PP_project/user_service/config"
	userproducer "github.com/JoePeach762/PP_project/user_service/internal/producer/user"
	"github.com/JoePeach762/PP_project/user_service/internal/services/user"
	statsstorage "github.com/JoePeach762/PP_project/user_service/internal/storage/pgstorage/statsstorage"
	userstorage "github.com/JoePeach762/PP_project/user_service/internal/storage/pgstorage/userstorage"
)

func InitUserService(
	userStorage *userstorage.PGstorage,
	statsStorage *statsstorage.PGstorage,
	producer *userproducer.UserKafkaProducer,
	cfg *config.Config,
) *user.Service {
	return user.NewUserService(
		userStorage,
		statsStorage,
		producer,
		uint32(cfg.UserServiceSettings.MinNameLen),
		uint32(cfg.UserServiceSettings.MaxNameLen),
		uint32(cfg.UserServiceSettings.MinWeight),
		uint32(cfg.UserServiceSettings.MaxWeight),
	)
}
