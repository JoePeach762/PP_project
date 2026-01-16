package pgstorage

import "time"

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
