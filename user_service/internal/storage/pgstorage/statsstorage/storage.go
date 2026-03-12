package statsstorage

import (
	"context"
	"fmt"
	"time"

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
			%s INTEGER NOT NULL DEFAULT 0 CHECK (%s >= 0),
			%s INTEGER NOT NULL DEFAULT 0 CHECK (%s >= 0),
			%s INTEGER NOT NULL DEFAULT 0 CHECK (%s >= 0),
			%s INTEGER NOT NULL DEFAULT 0 CHECK (%s >= 0),
			%s DATE NOT NULL DEFAULT CURRENT_DATE,
			PRIMARY KEY (%s, %s),
			FOREIGN KEY (%s) REFERENCES %s(%s) ON DELETE CASCADE
		)`, statsTableName,
		statsUserIDColumnName,
		statsCaloriesColumnName, statsCaloriesColumnName,
		statsProteinsColumnName, statsProteinsColumnName,
		statsFatsColumnName, statsFatsColumnName,
		statsCarbsColumnName, statsCarbsColumnName,
		statsDateColumnName,
		statsUserIDColumnName, statsDateColumnName,
		statsUserIDColumnName, usersTableName, usersIDColumnName,
	)

	_, err := s.db.Exec(context.Background(), statsSQL)
	if err != nil {
		return errors.Wrap(err, "init stats table")
	}

	processedMealEventsSQL := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s VARCHAR(128) PRIMARY KEY,
			%s BIGINT NOT NULL,
			%s TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
		)`, processedMealEventsTableName,
		processedMealEventIDColumnName,
		processedMealUserIDColumnName,
		processedMealProcessedAtColumnName,
	)

	_, err = s.db.Exec(context.Background(), processedMealEventsSQL)
	if err != nil {
		return errors.Wrap(err, "init processed meal events table")
	}

	return nil
}

func normalizeDate(date time.Time) time.Time {
	if date.IsZero() {
		date = time.Now()
	}

	year, month, day := date.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, date.Location())
}
