package models

import "time"

type StudentInfo struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   uint64 `json:"age"`
}

type UserInfo struct {
	ID              uint64 `json:"id"`
	Name            string `json:"name"`             //John
	Email           string `json:"email"`            //John@gmail.com
	Sex             string `json:"sex"`              //Male
	Age             uint32 `json:"age"`              //24
	HeightCm        uint32 `json:"height_cm"`        //180
	WeightKg        uint32 `json:"weight_kg"`        //80
	TargetWeightKg  uint32 `json:"target_weight_kg"` //95
	CurrentCalories uint32 `json:"current_calories"`
	CurrentProteins uint32 `json:"current_proteins"`
	CurrentFats     uint32 `json:"current_fats"`
	CurrentCarbs    uint32 `json:"current_carbs"`
	TargetCalories  uint32 `json:"target_calories"`
	TargetProteins  uint32 `json:"target_proteins"`
	TargetFats      uint32 `json:"target_fats"`
	TargetCarbs     uint32 `json:"target_carbs"`
}

type MealInput struct {
	Name        string  `json:"name"`
	WeightGrams float32 `json:"weight_grams"`
	UserID      uint64  `json:"user_id"`
}

type OFFProduct struct {
	Name         string  `json:"name"`
	Calories100g float32 `json:"calories_100g"`
	Proteins100g float32 `json:"proteins_100g"`
	Fats100g     float32 `json:"fats_100g"`
	Carbs100g    float32 `json:"carbs_100g"`
}

type MealInfo struct {
	ID           uint64    `json:"id"`
	Name         string    `json:"name"`         //Grilled chicken
	WeightGrams  float32   `json:"weight_grams"` //200
	Calories100g float32   `json:"calories_100g"`
	Proteins100g float32   `json:"proteins_100g"`
	Fats100g     float32   `json:"fats_100g"`
	Carbs100g    float32   `json:"carbs_100g"`
	Date         time.Time `json:"date"`
}
