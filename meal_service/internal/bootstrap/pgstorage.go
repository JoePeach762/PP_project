package bootstrap

import (
	"fmt"

	"github.com/JoePeach762/PP_project/meal_service/config"
	"github.com/JoePeach762/PP_project/meal_service/internal/storage/pgstorage"
	"github.com/pkg/errors"
)

func InitPGStorage(cfg *config.Config) (*pgstorage.PGstorage, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	storage, err := pgstorage.NewPGStorage(connString)
	if err != nil {
		return nil, errors.Wrap(err, "ошибка инициализации PostgreSQL")
	}
	return storage, nil
}
