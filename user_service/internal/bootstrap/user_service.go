package bootstrap

import (
	"github.com/JoePeach762/PP_project/user_service/config"
	"github.com/JoePeach762/PP_project/user_service/internal/services/user"
	"github.com/JoePeach762/PP_project/user_service/internal/storage/pgstorage"
)

func InitUserService(storage *pgstorage.PGstorage, cfg *config.Config) *user.Service {
	return user.NewUserService(
		storage,
		uint32(cfg.UserServiceSettings.MinNameLen),
		uint32(cfg.UserServiceSettings.MaxNameLen),
		uint32(cfg.UserServiceSettings.MinWeight),
		uint32(cfg.UserServiceSettings.MaxWeight),
	)
}
