package bootstrap

import (
	"github.com/JoePeach762/PP_project/user_service/config"
	"github.com/JoePeach762/PP_project/user_service/internal/services/user"
	statsstorage "github.com/JoePeach762/PP_project/user_service/internal/storage/pgstorage/statsstorage"
	userstorage "github.com/JoePeach762/PP_project/user_service/internal/storage/pgstorage/userstorage"
)

func InitUserService(
	userStorage *userstorage.PGstorage,
	statsStorage *statsstorage.PGstorage,
	cfg *config.Config,
) *user.Service {
	return user.NewUserService(
		userStorage,
		statsStorage,
		uint32(cfg.UserServiceSettings.MinNameLen),
		uint32(cfg.UserServiceSettings.MaxNameLen),
		uint32(cfg.UserServiceSettings.MinWeight),
		uint32(cfg.UserServiceSettings.MaxWeight),
	)
}
