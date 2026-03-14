package userstorage

import (
	"context"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (storage *PGstorage) GetUsersByIds(ctx context.Context, ids []uint64) ([]*models.UserInfo, error) {
	query := storage.getUsersQuery(ids)
	queryText, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "не удалось сформировать запрос на получение пользователей")
	}
	rows, err := storage.db.Query(ctx, queryText, args...)
	if err != nil {
		return nil, errors.Wrap(err, "не удалось выполнить запрос на получение пользователей")
	}
	defer rows.Close()

	var users []*models.UserInfo
	for rows.Next() {
		var u models.UserInfo
		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Sex,
			&u.Age,
			&u.HeightCm,
			&u.WeightKg,
			&u.TargetWeightKg,
			&u.TargetCalories,
			&u.TargetProteins,
			&u.TargetFats,
			&u.TargetCarbs,
		); err != nil {
			return nil, errors.Wrap(err, "не удалось прочитать строку пользователя")
		}
		users = append(users, &u)
	}
	return users, nil
}

func (storage *PGstorage) getUsersQuery(IDs []uint64) squirrel.Sqlizer {
	q := squirrel.
		Select(
			userIDColumnName,
			userNameColumnName,
			userEmailColumnName,
			userSexColumnName,
			userAgeColumnName,
			userHeightCmColumnName,
			userWeightKgColumnName,
			userTargetWeightKgColumnName,
			userTargetCaloriesColumnName,
			userTargetProteinsColumnName,
			userTargetFatsColumnName,
			userTargetCarbsColumnName,
		).
		From(userTableName).
		Where(squirrel.Eq{userIDColumnName: IDs}).
		PlaceholderFormat(squirrel.Dollar)
	return q
}
