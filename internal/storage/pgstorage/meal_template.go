package pgstorage

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MealTemplateStorage interface {
	Create(ctx context.Context, mt *MealTemplate) error
	GetByID(ctx context.Context, id string) (*MealTemplate, error)
	ListAll(ctx context.Context) ([]*MealTemplate, error)
	Update(ctx context.Context, mt *MealTemplate) error
}

type mealTemplateStorage struct {
	pool *pgxpool.Pool
}

func (s *mealTemplateStorage) Create(ctx context.Context, mt *MealTemplate) error {
	if mt.ID == "" {
		mt.ID = uuid.NewString()
	}
	if mt.CreatedAt.IsZero() {
		mt.CreatedAt = time.Now()
	}
	mt.UpdatedAt = mt.CreatedAt

	_, err := s.pool.Exec(ctx,
		`INSERT INTO meal_templates (
			id, name, description, tags, kcal, proteins,
			carbohydrates, fat, servings, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		mt.ID, mt.Name, mt.Description, StringSlice(mt.Tags),
		mt.Kcal, mt.Proteins, mt.Carbohydrates, mt.Fat, mt.Servings,
		mt.CreatedAt, mt.UpdatedAt,
	)
	return err
}

func (s *mealTemplateStorage) GetByID(ctx context.Context, id string) (*MealTemplate, error) {
	var mt MealTemplate
	var tags StringSlice
	err := s.pool.QueryRow(ctx,
		`SELECT id, name, description, tags, kcal, proteins,
		        carbohydrates, fat, servings, created_at, updated_at
		 FROM meal_templates WHERE id = $1`,
		id,
	).Scan(
		&mt.ID, &mt.Name, &mt.Description, &tags,
		&mt.Kcal, &mt.Proteins, &mt.Carbohydrates, &mt.Fat, &mt.Servings,
		&mt.CreatedAt, &mt.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	mt.Tags = []string(tags)
	return &mt, nil
}

func (s *mealTemplateStorage) ListAll(ctx context.Context) ([]*MealTemplate, error) {
	rows, err := s.pool.Query(ctx,
		`SELECT id, name, description, tags, kcal, proteins,
		        carbohydrates, fat, servings, created_at, updated_at
		 FROM meal_templates`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var templates []*MealTemplate
	for rows.Next() {
		var mt MealTemplate
		var tags StringSlice
		err := rows.Scan(
			&mt.ID, &mt.Name, &mt.Description, &tags,
			&mt.Kcal, &mt.Proteins, &mt.Carbohydrates, &mt.Fat, &mt.Servings,
			&mt.CreatedAt, &mt.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		mt.Tags = []string(tags)
		templates = append(templates, &mt)
	}
	return templates, nil
}

func (s *mealTemplateStorage) Update(ctx context.Context, mt *MealTemplate) error {
	mt.UpdatedAt = time.Now()
	_, err := s.pool.Exec(ctx,
		`UPDATE meal_templates SET
			name = $2, description = $3, tags = $4, kcal = $5,
			proteins = $6, carbohydrates = $7, fat = $8,
			servings = $9, updated_at = $10
		 WHERE id = $1`,
		mt.ID, mt.Name, mt.Description, StringSlice(mt.Tags),
		mt.Kcal, mt.Proteins, mt.Carbohydrates, mt.Fat, mt.Servings, mt.UpdatedAt,
	)
	return err
}
