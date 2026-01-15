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

	_, err := s.db.Exec(context.Background(), userSQL)
	if err != nil {
		return errors.Wrap(err, "init users table")
	}

	mealSQL := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s SERIAL PRIMARY KEY,
			%s BIGINT NOT NULL REFERENCES %s(%s) ON DELETE CASCADE,
			%s VARCHAR(255) NOT NULL,
			%s FLOAT4 NOT NULL CHECK (%s > 0),
			%s FLOAT4 NOT NULL,
			%s FLOAT4 NOT NULL,
			%s FLOAT4 NOT NULL,
			%s FLOAT4 NOT NULL,
			%s TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
		)`, mealTableName,
		mealIDColumnName,
		mealUserIDcolumnName, userTableName, userIDColumnName,
		mealNameColumnName,
		mealWeightGramsColumnName, mealWeightGramsColumnName,
		mealCalories100gColumnName,
		mealProteins100gColumnName,
		mealFats100gColumnName,
		mealCarbs100gColumnName,
		mealDateColumnName,
	)

	_, err = s.db.Exec(context.Background(), mealSQL)
	if err != nil {
		return errors.Wrap(err, "init meal_info table")
	}

	indexSQL := fmt.Sprintf(`
	CREATE INDEX IF NOT EXISTS idx_meal_info_user_date 
	ON %s (%s, %s)`,
		mealTableName,
		mealUserIDcolumnName,
		mealDateColumnName,
	)

	_, err = s.db.Exec(context.Background(), indexSQL)
	if err != nil {
		return errors.Wrap(err, "create meal_info index")
	}

	return nil
}
