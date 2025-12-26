package pgstorage

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// PgStorage — фабрика для получения интерфейсов хранилищ
type PgStorage interface {
	User() UserStorage
	UserPreferences() UserPreferencesStorage
	Product() ProductStorage
	MealLog() MealLogStorage
	DailySummary() DailySummaryStorage
	MealTemplate() MealTemplateStorage
	MealPlan() MealPlanStorage
	MealPlanDay() MealPlanDayStorage
	MealPlanEntry() MealPlanEntryStorage
}

type pgStorage struct {
	pool *pgxpool.Pool
}

func NewPgStorage(pool *pgxpool.Pool) PgStorage {
	return &pgStorage{pool: pool}
}

func (s *pgStorage) User() UserStorage {
	return &userStorage{pool: s.pool}
}

func (s *pgStorage) UserPreferences() UserPreferencesStorage {
	return &userPreferencesStorage{pool: s.pool}
}

func (s *pgStorage) Product() ProductStorage {
	return &productStorage{pool: s.pool}
}

func (s *pgStorage) MealLog() MealLogStorage {
	return &mealLogStorage{pool: s.pool}
}

func (s *pgStorage) DailySummary() DailySummaryStorage {
	return &dailySummaryStorage{pool: s.pool}
}

func (s *pgStorage) MealTemplate() MealTemplateStorage {
	return &mealTemplateStorage{pool: s.pool}
}

func (s *pgStorage) MealPlan() MealPlanStorage {
	return &mealPlanStorage{pool: s.pool}
}

func (s *pgStorage) MealPlanDay() MealPlanDayStorage {
	return &mealPlanDayStorage{pool: s.pool}
}

func (s *pgStorage) MealPlanEntry() MealPlanEntryStorage {
	return &mealPlanEntryStorage{pool: s.pool}
}
