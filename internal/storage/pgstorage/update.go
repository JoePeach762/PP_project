package pgstorage

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (s *PGstorage) UpdateUser(ctx context.Context, id uint64, info models.UserInfo) error {
	query := squirrel.Update(userTableName).
		Set(userNameColumnName, info.Name).
		Set(userEmailColumnName, info.Email).
		Set(userSexColumnName, info.Sex).
		Set(userAgeColumnName, info.Age).
		Set(userHeightCmColumnName, info.HeightCm).
		Set(userWeightKgColumnName, info.WeightKg).
		Set(userTargetWeightKgColumnName, info.TargetWeightKg).
		// Set(userCurrentCaloriesColumnName, info.CurrentCalories).	отдельно при получении события из кафки
		// Set(userCurrentProteinsColumnName, info.CurrentProteins).
		// Set(userCurrentFatsColumnName, info.CurrentFats).
		// Set(userCurrentCarbsColumnName, info.CurrentCarbs).
		Set(userTargetCaloriesColumnName, info.TargetCalories).
		Set(userTargetProteinsColumnName, info.TargetProteins).
		Set(userTargetFatsColumnName, info.TargetFats).
		Set(userTargetCarbsColumnName, info.TargetCarbs).
		Where(squirrel.Eq{userIDColumnName: id}).
		PlaceholderFormat(squirrel.Dollar)

	queryText, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "generate UPDATE query")
	}

	result, err := s.db.Exec(ctx, queryText, args...)
	if err != nil {
		return errors.Wrap(err, "execute UPDATE query")
	}

	if result.RowsAffected() == 0 {
		return errors.New("user not found")
	}

	return nil
}
