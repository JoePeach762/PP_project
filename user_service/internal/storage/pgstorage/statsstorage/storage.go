package statsstorage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type PGstorage struct {
	db *pgxpool.Pool
}

func NewPGStorage(connString string) (*PGstorage, error) {

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, errors.Wrap(err, "ошибка парсинга конфига")
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, errors.Wrap(err, "ошибка подключения")
	}
	storage := &PGstorage{
		db: db,
	}
	err = storage.initTables()
	if err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *PGstorage) initTables() error {
	statsSQL := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s BIGINT NOT NULL,
			%s SMALLINT DEFAULT 0 CHECK (%s > 0),
			%s SMALLINT DEFAULT 0 CHECK (%s > 0),
			%s SMALLINT DEFAULT 0 CHECK (%s > 0),
			%s SMALLINT DEFAULT 0 CHECK (%s > 0),
			%s TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
		)`, statsTableName,
		statsUserIDColumnName,
		statsCaloriesColumnName, statsCaloriesColumnName,
		statsProteinsColumnName, statsProteinsColumnName,
		statsFatsColumnName, statsFatsColumnName,
		statsCarbsColumnName, statsCarbsColumnName,
		statsDateColumnName,
	)

	_, err := s.db.Exec(context.Background(), statsSQL)
	if err != nil {
		return errors.Wrap(err, "init stats table")
	}

	return nil
}
