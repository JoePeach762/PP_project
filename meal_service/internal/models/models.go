package models

type MealInput struct {
	Name        string  `json:"name"`
	WeightGrams float32 `json:"weight_grams"`
	UserID      uint64  `json:"user_id"`
}

type MealTemplate struct {
	Name         string  `json:"name"`
	Calories100g float32 `json:"calories_100g"`
	Proteins100g float32 `json:"proteins_100g"`
	Fats100g     float32 `json:"fats_100g"`
	Carbs100g    float32 `json:"carbs_100g"`
}
