package bootstrap

import (
	"github.com/JoePeach762/PP_project/internal/storage/pgstorage"
	"github.com/pkg/errors"
)

func NewPGStorage(connString string) (*pgstorage.PGstorage, error) {
	storage, err := pgstorage.NewPGStorage(connString)
	if err != nil {
		return nil, errors.Wrap(err, "ошибка инициализации PostgreSQL")
	}
	return storage, nil
}
