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
	Age             uint8  `db:"age"`
	WeightKg        uint8  `db:"weightKg"`
	TargetWeightKg  uint8  `db:"targetWeightKg"`
	CurrentCalories uint16 `db:"currentCalories"`
	CurrentProteins uint16 `db:"currentProteins"`
	CurrentFats     uint16 `db:"currentFats"`
	CurrentCarbs    uint16 `db:"currentCarbs"`
}

type MealInfo struct {
	ID           uint64    `db:"id"`
	Name         string    `db:"name"`
	WeightGramms float32   `db:"weightGramms"`
	Calories100g float32   `db:"calories100g"`
	Proteins100g float32   `db:"proteins100g"`
	Fats100g     float32   `db:"fats100g"`
	Carbs100g    float32   `db:"carbs100g"`
	Date         time.Time `db:"date"`
}

const (
	studentTableName = "studentsInfo"

	studentIDColumnName    = "id"
	studentNameColumnName  = "name"
	studentEmailColumnName = "email"
	studentAgeColumnName   = "age"
)

const (
	userTableName = "usersInfo"

	userIDColumnName              = "id"
	userNameColumnName            = "name"
	userEmailColumnName           = "email"
	userSexColumnName             = "sex"
	userAgeColumnName             = "age"
	userWeightKgColumnName        = "weightKg"
	userTargetWeightKgColumnName  = "targetWeightKg"
	userCurrentCaloriesColumnName = "currentCalories"
	userCurrentProteinsColumnName = "currentProteins"
	userCurrentFatsColumnName     = "currentFats"
	userCurrentCarbsColumnName    = "currentCarbs"
)

// const (
// 	mealTableName = "mealInfo"

// 	mealIDColumnName = "id"
// 	mealNameColumnName = "name"
// 	mealWeightGrammsColumnName = "weightGramms"
// 	mealCalories100gColumnName = "calories100g"
// 	mealProteins100gColumnName = "proteins100g"
// 	mealFats100gColumnName = "fats100g"
// 	mealCarbs100gColumnName = "carbs100g"
// 	mealDateColumnName = "date"
// )
