package pgstorage

import "time"

type StudentInfo struct {
	ID    uint64 `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
	Age   uint64 `db:"age"`
}

type UserInfo struct {
	ID              uint64 `db:"id"`
	Name            string `db:"name"`
	Email           string `db:"email"`
	Sex             string `db:"sex"`
	Age             uint32 `db:"age"`
	HeightCm        uint32 `db:"height_cm"`
	WeightKg        uint32 `db:"weight_kg"`
	TargetWeightKg  uint32 `db:"target_weight_kg"`
	CurrentCalories uint32 `db:"current_calories"`
	CurrentProteins uint32 `db:"current_proteins"`
	CurrentFats     uint32 `db:"current_fats"`
	CurrentCarbs    uint32 `db:"current_carbs"`
	TargetCalories  uint32 `db:"target_calories"`
	TargetProteins  uint32 `db:"target_proteins"`
	TargetFats      uint32 `db:"target_fats"`
	TargetCarbs     uint32 `db:"target_carbs"`
}

type MealInfo struct {
	ID           uint64    `db:"id"`
	Name         string    `db:"name"`
	WeightGramms float32   `db:"weight_grams"`
	Calories100g float32   `db:"calories_100g"`
	Proteins100g float32   `db:"proteins_100g"`
	Fats100g     float32   `db:"fats_100g"`
	Carbs100g    float32   `db:"carbs_100g"`
	Date         time.Time `db:"date"`
}

const (
	studentTableName = "students_info"

	studentIDColumnName    = "id"
	studentNameColumnName  = "name"
	studentEmailColumnName = "email"
	studentAgeColumnName   = "age"
)

const (
	userTableName = "users_info"

	userIDColumnName              = "id"
	userNameColumnName            = "name"
	userEmailColumnName           = "email"
	userSexColumnName             = "sex"
	userAgeColumnName             = "age"
	userHeightCmColumnName        = "height_cm"
	userWeightKgColumnName        = "weight_kg"
	userTargetWeightKgColumnName  = "target_weight_kg"
	userCurrentCaloriesColumnName = "current_calories"
	userCurrentProteinsColumnName = "current_proteins"
	userCurrentFatsColumnName     = "current_fats"
	userCurrentCarbsColumnName    = "current_carbs"
	userTargetCaloriesColumnName  = "target_calories"
	userTargetProteinsColumnName  = "target_proteins"
	userTargetFatsColumnName      = "target_fats"
	userTargetCarbsColumnName     = "target_carbs"
)
