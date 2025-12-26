package pgstorage

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MealPlanStorage interface {
	Create(ctx context.Context, mp *MealPlan) error
	GetByID(ctx context.Context, id string) (*MealPlan, error)
	GetByUserIDAndWeekStart(ctx context.Context, userID string, weekStart time.Time) (*MealPlan, error)
	Update(ctx context.Context, mp *MealPlan) error
}

type mealPlanStorage struct {
	pool *pgxpool.Pool
}

func (s *mealPlanStorage) Create(ctx context.Context, mp *MealPlan) error {
	if mp.ID == "" {
		mp.ID = uuid.NewString()
	}
	if mp.CreatedAt.IsZero() {
		mp.CreatedAt = time.Now()
	}
	mp.UpdatedAt = mp.CreatedAt

	_, err := s.pool.Exec(ctx,
		`INSERT INTO meal_plans (id, user_id, week_start, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5)`,
		mp.ID, mp.UserID, mp.WeekStart, mp.CreatedAt, mp.UpdatedAt,
	)
	return err
}

func (s *mealPlanStorage) GetByID(ctx context.Context, id string) (*MealPlan, error) {
	var mp MealPlan
	err := s.pool.QueryRow(ctx,
		`SELECT id, user_id, week_start, created_at, updated_at
		 FROM meal_plans WHERE id = $1`,
		id,
	).Scan(&mp.ID, &mp.UserID, &mp.WeekStart, &mp.CreatedAt, &mp.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &mp, nil
}

func (s *mealPlanStorage) GetByUserIDAndWeekStart(ctx context.Context, userID string, weekStart time.Time) (*MealPlan, error) {
	var mp MealPlan
	err := s.pool.QueryRow(ctx,
		`SELECT id, user_id, week_start, created_at, updated_at
		 FROM meal_plans
		 WHERE user_id = $1 AND week_start = $2`,
		userID, weekStart,
	).Scan(&mp.ID, &mp.UserID, &mp.WeekStart, &mp.CreatedAt, &mp.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &mp, nil
}

func (s *mealPlanStorage) Update(ctx context.Context, mp *MealPlan) error {
	mp.UpdatedAt = time.Now()
	_, err := s.pool.Exec(ctx,
		`UPDATE meal_plans SET week_start = $2, updated_at = $3 WHERE id = $1`,
		mp.ID, mp.WeekStart, mp.UpdatedAt,
	)
	return err
}
