package userstorage

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
	userSQL := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s SERIAL PRIMARY KEY,
			%s VARCHAR(100) NOT NULL,
			%s VARCHAR(255) UNIQUE NOT NULL,
			%s VARCHAR(10) NOT NULL CHECK (%s IN ('male', 'female')),
			%s SMALLINT NOT NULL CHECK (%s > 0 AND %s < 100),
			%s SMALLINT NOT NULL CHECK (%s > 0 AND %s < 300),
			%s SMALLINT NOT NULL CHECK (%s > 0),
			%s SMALLINT NOT NULL CHECK (%s > 0),
			%s SMALLINT DEFAULT 0 CHECK (%s >= 0),
			%s SMALLINT DEFAULT 0 CHECK (%s >= 0),
			%s SMALLINT DEFAULT 0 CHECK (%s >= 0),
			%s SMALLINT DEFAULT 0 CHECK (%s >= 0)
		)`, userTableName,
		userIDColumnName,
		userNameColumnName,
		userEmailColumnName,
		userSexColumnName, userSexColumnName,
		userAgeColumnName, userAgeColumnName, userAgeColumnName,
		userHeightCmColumnName, userHeightCmColumnName, userHeightCmColumnName,
		userWeightKgColumnName, userWeightKgColumnName,
		userTargetWeightKgColumnName, userTargetWeightKgColumnName,
		userTargetCaloriesColumnName, userTargetCaloriesColumnName,
		userTargetProteinsColumnName, userTargetProteinsColumnName,
		userTargetFatsColumnName, userTargetFatsColumnName,
		userTargetCarbsColumnName, userTargetCarbsColumnName,
	)

	_, err := s.db.Exec(context.Background(), userSQL)
	if err != nil {
		return errors.Wrap(err, "init users table")
	}

	outboxSQL := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s BIGSERIAL PRIMARY KEY,
			%s BYTEA NOT NULL,
			%s TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
			%s TIMESTAMP WITH TIME ZONE
		)`, outboxTableName,
		outboxIDColumnName,
		outboxPayloadColumnName,
		outboxCreatedAtColumnName,
		outboxPublishedAtColumnName,
	)

	_, err = s.db.Exec(context.Background(), outboxSQL)
	if err != nil {
		return errors.Wrap(err, "init user outbox table")
	}

	outboxIndexSQL := fmt.Sprintf(`
		CREATE INDEX IF NOT EXISTS idx_%s_unpublished
		ON %s (%s, %s)`, outboxTableName,
		outboxTableName,
		outboxPublishedAtColumnName,
		outboxIDColumnName,
	)

	_, err = s.db.Exec(context.Background(), outboxIndexSQL)
	if err != nil {
		return errors.Wrap(err, "create user outbox index")
	}

	return nil
}
