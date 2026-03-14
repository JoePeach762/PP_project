package mealstorage

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
	mealSQL := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s SERIAL PRIMARY KEY,
			%s BIGINT NOT NULL,
			%s VARCHAR(255) NOT NULL,
			%s FLOAT4 NOT NULL CHECK (%s > 0),
			%s FLOAT4 NOT NULL,
			%s FLOAT4 NOT NULL,
			%s FLOAT4 NOT NULL,
			%s FLOAT4 NOT NULL,
			%s TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
		)`, mealTableName,
		mealIDColumnName,
		mealUserIDcolumnName,
		mealNameColumnName,
		mealWeightGramsColumnName, mealWeightGramsColumnName,
		mealCalories100gColumnName,
		mealProteins100gColumnName,
		mealFats100gColumnName,
		mealCarbs100gColumnName,
		mealDateColumnName,
	)

	_, err := s.db.Exec(context.Background(), mealSQL)
	if err != nil {
		return errors.Wrap(err, "не удалось инициализировать таблицу приемов пищи")
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
		return errors.Wrap(err, "не удалось создать индекс таблицы приемов пищи")
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
		return errors.Wrap(err, "не удалось инициализировать таблицу meal outbox")
	}

	outboxIndexSQL := fmt.Sprintf(`
	CREATE INDEX IF NOT EXISTS idx_%s_unpublished
	ON %s (%s, %s)`,
		outboxTableName,
		outboxTableName,
		outboxPublishedAtColumnName,
		outboxIDColumnName,
	)

	_, err = s.db.Exec(context.Background(), outboxIndexSQL)
	if err != nil {
		return errors.Wrap(err, "не удалось создать индекс для meal outbox")
	}

	return nil
}
