package models

import (
	"time"
)

type User struct {
	ID        string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserPreferences struct {
	UserID    string
	Interests []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	Code          string
	Name          string
	Brands        string
	EnergyKcal    float64
	Proteins      float64
	Carbohydrates float64
	Fat           float64
	Salt          float64
	NutriScore    string
	EcoScore      string
	FetchedAt     time.Time
}

type MealLog struct {
	ID            string
	UserID        string
	ProductCode   string
	ProductName   string
	Grams         float64
	Kcal          float64
	Proteins      float64
	Carbohydrates float64
	Fat           float64
	LoggedAt      time.Time
	CreatedAt     time.Time
}

type DailySummary struct {
	UserID        string
	Date          time.Time
	TotalKcal     float64
	TotalProteins float64
	TotalCarbs    float64
	TotalFat      float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type MealTemplate struct {
	ID            string
	Name          string
	Description   string
	Tags          []string
	Kcal          float64
	Proteins      float64
	Carbohydrates float64
	Fat           float64
	Servings      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type MealPlan struct {
	ID        string
	UserID    string
	WeekStart time.Time
	Days      []MealPlanDay
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MealPlanDay struct {
	Date  time.Time
	Meals []MealPlanEntry
}

type MealPlanEntry struct {
	MealType   string
	TemplateID string
	Servings   int
}

type MealEvent struct {
	UserID      string
	ProductCode string
	Grams       float64
	LoggedAt    time.Time
}

type DailyResetEvent struct {
	TriggeredAt time.Time
}

type StudentInfo struct {
	ID    uint64
	Name  string
	Email string
	Age   uint64
}
