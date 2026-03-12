package userstorage

type UserInfo struct {
	ID             uint64 `db:"id"`
	Name           string `db:"name"`
	Email          string `db:"email"`
	Sex            string `db:"sex"`
	Age            uint32 `db:"age"`
	HeightCm       uint32 `db:"height_cm"`
	WeightKg       uint32 `db:"weight_kg"`
	TargetWeightKg uint32 `db:"target_weight_kg"`
	TargetCalories uint32 `db:"target_calories"`
	TargetProteins uint32 `db:"target_proteins"`
	TargetFats     uint32 `db:"target_fats"`
	TargetCarbs    uint32 `db:"target_carbs"`
}

const (
	userTableName = "users_info"

	userIDColumnName             = "id"
	userNameColumnName           = "name"
	userEmailColumnName          = "email"
	userSexColumnName            = "sex"
	userAgeColumnName            = "age"
	userHeightCmColumnName       = "height_cm"
	userWeightKgColumnName       = "weight_kg"
	userTargetWeightKgColumnName = "target_weight_kg"
	userTargetCaloriesColumnName = "target_calories"
	userTargetProteinsColumnName = "target_proteins"
	userTargetFatsColumnName     = "target_fats"
	userTargetCarbsColumnName    = "target_carbs"
)
