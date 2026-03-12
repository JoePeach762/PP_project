package models

import "time"

type UserInput struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Sex            string `json:"sex"`
	Age            uint32 `json:"age"`
	HeightCm       uint32 `json:"height_cm"`
	WeightKg       uint32 `json:"weight_kg"`
	TargetWeightKg uint32 `json:"target_weight_kg"`
}

type UserInfo struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Sex            string `json:"sex"`
	Age            uint32 `json:"age"`
	HeightCm       uint32 `json:"height_cm"`
	WeightKg       uint32 `json:"weight_kg"`
	TargetWeightKg uint32 `json:"target_weight_kg"`
	TargetCalories uint32 `json:"target_calories"`
	TargetProteins uint32 `json:"target_proteins"`
	TargetFats     uint32 `json:"target_fats"`
	TargetCarbs    uint32 `json:"target_carbs"`
}

type UserStats struct {
	UserID          uint64    `json:"user_id"`
	CurrentCalories uint32    `json:"current_calories"`
	CurrentProteins uint32    `json:"current_proteins"`
	CurrentFats     uint32    `json:"current_fats"`
	CurrentCarbs    uint32    `json:"current_carbs"`
	Date            time.Time `json:"date"`
}

type UserWithStats struct {
	ID              uint64 `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Sex             string `json:"sex"`
	Age             uint32 `json:"age"`
	HeightCm        uint32 `json:"height_cm"`
	WeightKg        uint32 `json:"weight_kg"`
	TargetWeightKg  uint32 `json:"target_weight_kg"`
	CurrentCalories uint32 `json:"current_calories"`
	CurrentProteins uint32 `json:"current_proteins"`
	CurrentFats     uint32 `json:"current_fats"`
	CurrentCarbs    uint32 `json:"current_carbs"`
	TargetCalories  uint32 `json:"target_calories"`
	TargetProteins  uint32 `json:"target_proteins"`
	TargetFats      uint32 `json:"target_fats"`
	TargetCarbs     uint32 `json:"target_carbs"`
}

type MealInfo struct {
	EventID      string    `json:"event_id"`
	ID           uint64    `json:"id"`
	UserId       uint64    `json:"user_id"`
	Name         string    `json:"name"`
	WeightGrams  float32   `json:"weight_grams"`
	Calories100g float32   `json:"calories_100g"`
	Proteins100g float32   `json:"proteins_100g"`
	Fats100g     float32   `json:"fats_100g"`
	Carbs100g    float32   `json:"carbs_100g"`
	Date         time.Time `json:"date"`
}

func NewUserInfoFromInput(in *UserInput) *UserInfo {
	return &UserInfo{
		Name:           in.Name,
		Email:          in.Email,
		Sex:            in.Sex,
		Age:            in.Age,
		HeightCm:       in.HeightCm,
		WeightKg:       in.WeightKg,
		TargetWeightKg: in.TargetWeightKg,
	}
}

func NewUserWithStats(info *UserInfo, stats *UserStats) *UserWithStats {
	userWithStats := &UserWithStats{
		ID:             info.ID,
		Name:           info.Name,
		Email:          info.Email,
		Sex:            info.Sex,
		Age:            info.Age,
		HeightCm:       info.HeightCm,
		WeightKg:       info.WeightKg,
		TargetWeightKg: info.TargetWeightKg,
		TargetCalories: info.TargetCalories,
		TargetProteins: info.TargetProteins,
		TargetFats:     info.TargetFats,
		TargetCarbs:    info.TargetCarbs,
	}
	if stats == nil {
		return userWithStats
	}

	userWithStats.CurrentCalories = stats.CurrentCalories
	userWithStats.CurrentProteins = stats.CurrentProteins
	userWithStats.CurrentFats = stats.CurrentFats
	userWithStats.CurrentCarbs = stats.CurrentCarbs

	return userWithStats
}
