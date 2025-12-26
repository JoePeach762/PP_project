package pgstorage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductStorage interface {
	Create(ctx context.Context, p *Product) error
	GetByCode(ctx context.Context, code string) (*Product, error)
	Update(ctx context.Context, p *Product) error
}

type productStorage struct {
	pool *pgxpool.Pool
}

func (s *productStorage) Create(ctx context.Context, p *Product) error {
	if p.FetchedAt.IsZero() {
		p.FetchedAt = time.Now()
	}

	_, err := s.pool.Exec(ctx,
		`INSERT INTO products (
			code, name, brands, energy_kcal, proteins, carbohydrates,
			fat, salt, nutri_score, eco_score, fetched_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (code) DO NOTHING`,
		p.Code, p.Name, p.Brands, p.EnergyKcal, p.Proteins,
		p.Carbohydrates, p.Fat, p.Salt, p.NutriScore, p.EcoScore, p.FetchedAt,
	)
	return err
}

func (s *productStorage) GetByCode(ctx context.Context, code string) (*Product, error) {
	var p Product
	err := s.pool.QueryRow(ctx,
		`SELECT code, name, brands, energy_kcal, proteins, carbohydrates,
		        fat, salt, nutri_score, eco_score, fetched_at
		 FROM products WHERE code = $1`,
		code,
	).Scan(
		&p.Code, &p.Name, &p.Brands, &p.EnergyKcal, &p.Proteins,
		&p.Carbohydrates, &p.Fat, &p.Salt, &p.NutriScore, &p.EcoScore, &p.FetchedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (s *productStorage) Update(ctx context.Context, p *Product) error {
	p.FetchedAt = time.Now()
	_, err := s.pool.Exec(ctx,
		`UPDATE products SET
			name = $2, brands = $3, energy_kcal = $4, proteins = $5,
			carbohydrates = $6, fat = $7, salt = $8, nutri_score = $9,
			eco_score = $10, fetched_at = $11
		 WHERE code = $1`,
		p.Code, p.Name, p.Brands, p.EnergyKcal, p.Proteins,
		p.Carbohydrates, p.Fat, p.Salt, p.NutriScore, p.EcoScore, p.FetchedAt,
	)
	return err
}
