package pgstorage

import "time"

// StudentInfo оставлен как legacy-пример — можно удалить позже
type StudentInfo struct {
	ID    uint64 `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
	Age   uint64 `db:"age"`
}

type User struct {
	ID        string    `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserPreferences struct {
	UserID    string    `db:"user_id"`
	Interests []string  `db:"interests"` // JSONB в PostgreSQL
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Product struct {
	Code          string    `db:"code"`
	Name          string    `db:"name"`
	Brands        string    `db:"brands"`
	EnergyKcal    float64   `db:"energy_kcal"`
	Proteins      float64   `db:"proteins"`
	Carbohydrates float64   `db:"carbohydrates"`
	Fat           float64   `db:"fat"`
	Salt          float64   `db:"salt"`
	NutriScore    string    `db:"nutri_score"`
	EcoScore      string    `db:"eco_score"`
	FetchedAt     time.Time `db:"fetched_at"`
}

type MealLog struct {
	ID            string    `db:"id"`
	UserID        string    `db:"user_id"`
	ProductCode   string    `db:"product_code"`
	ProductName   string    `db:"product_name"`
	Grams         float64   `db:"grams"`
	Kcal          float64   `db:"kcal"`
	Proteins      float64   `db:"proteins"`
	Carbohydrates float64   `db:"carbohydrates"`
	Fat           float64   `db:"fat"`
	LoggedAt      time.Time `db:"logged_at"`
	CreatedAt     time.Time `db:"created_at"`
}

type DailySummary struct {
	UserID        string    `db:"user_id"`
	Date          time.Time `db:"date"`
	TotalKcal     float64   `db:"total_kcal"`
	TotalProteins float64   `db:"total_proteins"`
	TotalCarbs    float64   `db:"total_carbs"`
	TotalFat      float64   `db:"total_fat"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type MealTemplate struct {
	ID            string    `db:"id"`
	Name          string    `db:"name"`
	Description   string    `db:"description"`
	Tags          []string  `db:"tags"` // JSONB
	Kcal          float64   `db:"kcal"`
	Proteins      float64   `db:"proteins"`
	Carbohydrates float64   `db:"carbohydrates"`
	Fat           float64   `db:"fat"`
	Servings      int       `db:"servings"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type MealPlan struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	WeekStart time.Time `db:"week_start"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type MealPlanDay struct {
	ID     string    `db:"id"`
	PlanID string    `db:"plan_id"`
	Date   time.Time `db:"date"`
}

type MealPlanEntry struct {
	ID         string `db:"id"`
	DayID      string `db:"day_id"`
	MealType   string `db:"meal_type"`
	TemplateID string `db:"template_id"`
	Servings   int    `db:"servings"`
}
