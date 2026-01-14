package pgstorage

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
	sql := fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %v (
        %v SERIAL PRIMARY KEY,
        %v VARCHAR(100) NOT NULL,
        %v VARCHAR(255) UNIQUE NOT NULL,
        %v INT
    )`, studentTableName,
		studentIDColumnName,
		studentNameColumnName,
		studentEmailColumnName,
		studentAgeColumnName,
	)
	_, err := s.db.Exec(context.Background(), sql)
	if err != nil {
		return errors.Wrap(err, "init students tables")
	}

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
			%s SMALLINT DEFAULT 0 CHECK (%s >= 0),
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
		userCurrentCaloriesColumnName, userCurrentCaloriesColumnName,
		userCurrentProteinsColumnName, userCurrentProteinsColumnName,
		userCurrentFatsColumnName, userCurrentFatsColumnName,
		userCurrentCarbsColumnName, userCurrentCarbsColumnName,
		userTargetCaloriesColumnName, userTargetCaloriesColumnName,
		userTargetProteinsColumnName, userTargetProteinsColumnName,
		userTargetFatsColumnName, userTargetFatsColumnName,
		userTargetCarbsColumnName, userTargetCarbsColumnName,
	)

	_, err = s.db.Exec(context.Background(), userSQL)
	if err != nil {
		return errors.Wrap(err, "init users table")
	}

	return nil
}
