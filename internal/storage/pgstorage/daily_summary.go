package pgstorage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DailySummaryStorage interface {
	Upsert(ctx context.Context, summary *DailySummary) error
	GetByUserAndDate(ctx context.Context, userID string, date time.Time) (*DailySummary, error)
}

type dailySummaryStorage struct {
	pool *pgxpool.Pool
}

func (s *dailySummaryStorage) Upsert(ctx context.Context, ds *DailySummary) error {
	if ds.CreatedAt.IsZero() {
		ds.CreatedAt = time.Now()
	}
	ds.UpdatedAt = time.Now()

	_, err := s.pool.Exec(ctx,
		`INSERT INTO daily_summaries (
			user_id, date, total_kcal, total_proteins, total_carbs, total_fat, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (user_id, date)
		DO UPDATE SET 
			total_kcal = EXCLUDED.total_kcal,
			total_proteins = EXCLUDED.total_proteins,
			total_carbs = EXCLUDED.total_carbs,
			total_fat = EXCLUDED.total_fat,
			updated_at = EXCLUDED.updated_at`,
		ds.UserID, ds.Date, ds.TotalKcal, ds.TotalProteins, ds.TotalCarbs, ds.TotalFat,
		ds.CreatedAt, ds.UpdatedAt,
	)
	return err
}

func (s *dailySummaryStorage) GetByUserAndDate(ctx context.Context, userID string, date time.Time) (*DailySummary, error) {
	var ds DailySummary
	err := s.pool.QueryRow(ctx,
		`SELECT user_id, date, total_kcal, total_proteins, total_carbs, total_fat, created_at, updated_at
		 FROM daily_summaries
		 WHERE user_id = $1 AND date = $2`,
		userID, date,
	).Scan(
		&ds.UserID, &ds.Date, &ds.TotalKcal, &ds.TotalProteins, &ds.TotalCarbs, &ds.TotalFat,
		&ds.CreatedAt, &ds.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &ds, nil
}
