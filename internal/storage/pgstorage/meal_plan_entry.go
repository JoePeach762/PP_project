package pgstorage

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MealPlanEntryStorage interface {
	Create(ctx context.Context, mpe *MealPlanEntry) error
	GetByID(ctx context.Context, id string) (*MealPlanEntry, error)
	GetByDayID(ctx context.Context, dayID string) ([]*MealPlanEntry, error)
	Update(ctx context.Context, mpe *MealPlanEntry) error
}

type mealPlanEntryStorage struct {
	pool *pgxpool.Pool
}

func (s *mealPlanEntryStorage) Create(ctx context.Context, mpe *MealPlanEntry) error {
	if mpe.ID == "" {
		mpe.ID = uuid.NewString()
	}

	_, err := s.pool.Exec(ctx,
		`INSERT INTO meal_plan_entries (id, day_id, meal_type, template_id, servings)
		 VALUES ($1, $2, $3, $4, $5)`,
		mpe.ID, mpe.DayID, mpe.MealType, mpe.TemplateID, mpe.Servings,
	)
	return err
}

func (s *mealPlanEntryStorage) GetByID(ctx context.Context, id string) (*MealPlanEntry, error) {
	var mpe MealPlanEntry
	err := s.pool.QueryRow(ctx,
		`SELECT id, day_id, meal_type, template_id, servings
		 FROM meal_plan_entries WHERE id = $1`,
		id,
	).Scan(&mpe.ID, &mpe.DayID, &mpe.MealType, &mpe.TemplateID, &mpe.Servings)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &mpe, nil
}

func (s *mealPlanEntryStorage) GetByDayID(ctx context.Context, dayID string) ([]*MealPlanEntry, error) {
	rows, err := s.pool.Query(ctx,
		`SELECT id, day_id, meal_type, template_id, servings
		 FROM meal_plan_entries WHERE day_id = $1`,
		dayID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []*MealPlanEntry
	for rows.Next() {
		var e MealPlanEntry
		err := rows.Scan(&e.ID, &e.DayID, &e.MealType, &e.TemplateID, &e.Servings)
		if err != nil {
			return nil, err
		}
		entries = append(entries, &e)
	}
	return entries, nil
}

func (s *mealPlanEntryStorage) Update(ctx context.Context, mpe *MealPlanEntry) error {
	_, err := s.pool.Exec(ctx,
		`UPDATE meal_plan_entries SET
			day_id = $2, meal_type = $3, template_id = $4, servings = $5
		 WHERE id = $1`,
		mpe.ID, mpe.DayID, mpe.MealType, mpe.TemplateID, mpe.Servings,
	)
	return err
}
