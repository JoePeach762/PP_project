package pgstorage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserPreferencesStorage interface {
	Create(ctx context.Context, up *UserPreferences) error
	GetByUserID(ctx context.Context, userID string) (*UserPreferences, error)
	Update(ctx context.Context, up *UserPreferences) error
}

type userPreferencesStorage struct {
	pool *pgxpool.Pool
}

func (s *userPreferencesStorage) Create(ctx context.Context, up *UserPreferences) error {
	if up.CreatedAt.IsZero() {
		up.CreatedAt = time.Now()
	}
	up.UpdatedAt = up.CreatedAt

	_, err := s.pool.Exec(ctx,
		`INSERT INTO user_preferences (user_id, interests, created_at, updated_at)
		 VALUES ($1, $2, $3, $4)
		 ON CONFLICT (user_id) DO NOTHING`,
		up.UserID, StringSlice(up.Interests), up.CreatedAt, up.UpdatedAt,
	)
	return err
}

func (s *userPreferencesStorage) GetByUserID(ctx context.Context, userID string) (*UserPreferences, error) {
	var up UserPreferences
	var interests StringSlice
	err := s.pool.QueryRow(ctx,
		`SELECT user_id, interests, created_at, updated_at
		 FROM user_preferences WHERE user_id = $1`,
		userID,
	).Scan(&up.UserID, &interests, &up.CreatedAt, &up.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	up.Interests = []string(interests)
	return &up, nil
}

func (s *userPreferencesStorage) Update(ctx context.Context, up *UserPreferences) error {
	up.UpdatedAt = time.Now()
	_, err := s.pool.Exec(ctx,
		`UPDATE user_preferences SET interests = $2, updated_at = $3 WHERE user_id = $1`,
		up.UserID, StringSlice(up.Interests), up.UpdatedAt,
	)
	return err
}
