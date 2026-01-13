package models

import (
	"time"
)

type StudentInfo struct {
	ID    uint64
	Name  string
	Email string
	Age   uint64
}

type UserInfo struct {
	ID              uint64
	Name            string //John
	Email           string //John@gmail.com
	Sex             string //Male
	Age             uint8  //24
	WeightKg        uint8  //80
	TargetWeightKg  uint8  //95
	CurrentCalories uint16
	CurrentProteins uint16
	CurrentFats     uint16
	CurrentCarbs    uint16
}

type MealInfo struct {
	ID           uint64
	Name         string  //Grilled chicken
	WeightGramms float32 //200
	Calories100g float32
	Proteins100g float32
	Fats100g     float32
	Carbs100g    float32
	Date         time.Time
}
