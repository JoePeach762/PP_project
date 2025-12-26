package pgstorage

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MealPlanDayStorage interface {
	Create(ctx context.Context, mpd *MealPlanDay) error
	GetByID(ctx context.Context, id string) (*MealPlanDay, error)
	GetByPlanID(ctx context.Context, planID string) ([]*MealPlanDay, error)
	Update(ctx context.Context, mpd *MealPlanDay) error
}

type mealPlanDayStorage struct {
	pool *pgxpool.Pool
}

func (s *mealPlanDayStorage) Create(ctx context.Context, mpd *MealPlanDay) error {
	if mpd.ID == "" {
		mpd.ID = uuid.NewString()
	}

	_, err := s.pool.Exec(ctx,
		`INSERT INTO meal_plan_days (id, plan_id, date)
		 VALUES ($1, $2, $3)`,
		mpd.ID, mpd.PlanID, mpd.Date,
	)
	return err
}

func (s *mealPlanDayStorage) GetByID(ctx context.Context, id string) (*MealPlanDay, error) {
	var mpd MealPlanDay
	err := s.pool.QueryRow(ctx,
		`SELECT id, plan_id, date FROM meal_plan_days WHERE id = $1`,
		id,
	).Scan(&mpd.ID, &mpd.PlanID, &mpd.Date)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &mpd, nil
}

func (s *mealPlanDayStorage) GetByPlanID(ctx context.Context, planID string) ([]*MealPlanDay, error) {
	rows, err := s.pool.Query(ctx,
		`SELECT id, plan_id, date FROM meal_plan_days WHERE plan_id = $1 ORDER BY date`,
		planID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var days []*MealPlanDay
	for rows.Next() {
		var d MealPlanDay
		err := rows.Scan(&d.ID, &d.PlanID, &d.Date)
		if err != nil {
			return nil, err
		}
		days = append(days, &d)
	}
	return days, nil
}

func (s *mealPlanDayStorage) Update(ctx context.Context, mpd *MealPlanDay) error {
	_, err := s.pool.Exec(ctx,
		`UPDATE meal_plan_days SET plan_id = $2, date = $3 WHERE id = $1`,
		mpd.ID, mpd.PlanID, mpd.Date,
	)
	return err
}
