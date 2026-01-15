package pgstorage

import "time"

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
	UserId       uint64    `db:"user_id"`
	Name         string    `db:"name"`
	WeightGrams  float32   `db:"weight_grams"`
	Calories100g float32   `db:"calories_100g"`
	Proteins100g float32   `db:"proteins_100g"`
	Fats100g     float32   `db:"fats_100g"`
	Carbs100g    float32   `db:"carbs_100g"`
	Date         time.Time `db:"date"`
}

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

const (
	mealTableName              = "meal_info"
	mealIDColumnName           = "id"
	mealUserIDcolumnName       = "user_id"
	mealNameColumnName         = "name"
	mealWeightGramsColumnName  = "weight_grams"
	mealCalories100gColumnName = "calories_100g"
	mealProteins100gColumnName = "proteins_100g"
	mealFats100gColumnName     = "fats_100g"
	mealCarbs100gColumnName    = "carbs_100g"
	mealDateColumnName         = "date"
)
