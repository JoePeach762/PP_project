package bootstrap

import (
	"github.com/JoePeach762/PP_project/user_service/config"
	statsstorage "github.com/JoePeach762/PP_project/user_service/internal/storage/pgstorage/statsstorage"
	userstorage "github.com/JoePeach762/PP_project/user_service/internal/storage/pgstorage/userstorage"
	"github.com/pkg/errors"
)

func InitPGStorage(cfg *config.Config) (*userstorage.PGstorage, *statsstorage.PGstorage, error) {
	connString := cfg.Database.ConnString()

	userStorage, err := userstorage.NewPGStorage(connString)
	if err != nil {
		return nil, nil, errors.Wrap(err, "ошибка инициализации users PostgreSQL")
	}

	statsStorage, err := statsstorage.NewPGStorage(connString)
	if err != nil {
		return nil, nil, errors.Wrap(err, "ошибка инициализации stats PostgreSQL")
	}

	return userStorage, statsStorage, nil
}
