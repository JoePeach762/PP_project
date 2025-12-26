package pgstorage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStorage interface {
	Create(ctx context.Context, u *User) error
	GetByID(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, u *User) error
}

type userStorage struct {
	pool *pgxpool.Pool
}

func (s *userStorage) Create(ctx context.Context, u *User) error {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = u.CreatedAt

	_, err := s.pool.Exec(ctx,
		`INSERT INTO users (id, email, password, created_at, updated_at) 
		 VALUES ($1, $2, $3, $4, $5)`,
		u.ID, u.Email, u.Password, u.CreatedAt, u.UpdatedAt,
	)
	return err
}

func (s *userStorage) GetByID(ctx context.Context, id string) (*User, error) {
	var u User
	err := s.pool.QueryRow(ctx,
		`SELECT id, email, password, created_at, updated_at FROM users WHERE id = $1`,
		id,
	).Scan(&u.ID, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (s *userStorage) GetByEmail(ctx context.Context, email string) (*User, error) {
	var u User
	err := s.pool.QueryRow(ctx,
		`SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1`,
		email,
	).Scan(&u.ID, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (s *userStorage) Update(ctx context.Context, u *User) error {
	u.UpdatedAt = time.Now()
	_, err := s.pool.Exec(ctx,
		`UPDATE users SET email = $2, password = $3, updated_at = $4 WHERE id = $1`,
		u.ID, u.Email, u.Password, u.UpdatedAt,
	)
	return err
}
