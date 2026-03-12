package statsstorage

import "time"

type UserStats struct {
	UserID          uint64    `db:"user_id"`
	CurrentCalories uint32    `db:"calories"`
	CurrentProteins uint32    `db:"proteins"`
	CurrentFats     uint32    `db:"fats"`
	CurrentCarbs    uint32    `db:"carbs"`
	Date            time.Time `db:"date"`
}

const (
	statsTableName          = "daily_stats"
	statsUserIDColumnName   = "user_id"
	statsCaloriesColumnName = "calories"
	statsProteinsColumnName = "proteins"
	statsFatsColumnName     = "fats"
	statsCarbsColumnName    = "carbs"
	statsDateColumnName     = "date"
	usersTableName          = "users_info"
	usersIDColumnName       = "id"

	processedMealEventsTableName       = "processed_meal_events"
	processedMealEventIDColumnName     = "event_id"
	processedMealUserIDColumnName      = "user_id"
	processedMealProcessedAtColumnName = "processed_at"
)
