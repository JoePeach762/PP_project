package statsstorage

import "time"

type UserStats struct {
	UserID   uint64    `db:"user_id"`
	Calories uint32    `db:"calories"`
	Proteins uint32    `db:"proteins"`
	Fats     uint32    `db:"fats"`
	Carbs    uint32    `db:"carbs"`
	Date     time.Time `db:"date"`
}

const (
	statsTableName          = "daily_stats"
	statsUserIDColumnName   = "user_id"
	statsCaloriesColumnName = "calories"
	statsProteinsColumnName = "proteins"
	statsFatsColumnName     = "fats"
	statsCarbsColumnName    = "carbs"
	statsDateColumnName     = "date"
)
