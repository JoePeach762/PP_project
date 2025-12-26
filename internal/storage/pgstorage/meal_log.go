package pgstorage

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MealLogStorage interface {
	Create(ctx context.Context, log *MealLog) error
	GetByUserAndDate(ctx context.Context, userID string, date time.Time) ([]*MealLog, error)
}

type mealLogStorage struct {
	pool *pgxpool.Pool
}

func (s *mealLogStorage) Create(ctx context.Context, log *MealLog) error {
	if log.ID == "" {
		log.ID = uuid.NewString()
	}
	if log.CreatedAt.IsZero() {
		log.CreatedAt = time.Now()
	}

	_, err := s.pool.Exec(ctx,
		`INSERT INTO meal_logs (
			id, user_id, product_code, product_name, grams, 
			kcal, proteins, carbohydrates, fat, logged_at, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		log.ID, log.UserID, log.ProductCode, log.ProductName,
		log.Grams, log.Kcal, log.Proteins, log.Carbohydrates, log.Fat,
		log.LoggedAt, log.CreatedAt,
	)
	return err
}

func (s *mealLogStorage) GetByUserAndDate(ctx context.Context, userID string, date time.Time) ([]*MealLog, error) {
	start := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	end := start.Add(24 * time.Hour)

	rows, err := s.pool.Query(ctx,
		`SELECT id, user_id, product_code, product_name, grams, 
		        kcal, proteins, carbohydrates, fat, logged_at, created_at
		 FROM meal_logs 
		 WHERE user_id = $1 AND logged_at >= $2 AND logged_at < $3
		 ORDER BY logged_at`,
		userID, start, end,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []*MealLog
	for rows.Next() {
		var l MealLog
		err := rows.Scan(
			&l.ID, &l.UserID, &l.ProductCode, &l.ProductName,
			&l.Grams, &l.Kcal, &l.Proteins, &l.Carbohydrates, &l.Fat,
			&l.LoggedAt, &l.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		logs = append(logs, &l)
	}
	return logs, nil
}
