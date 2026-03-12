package userstorage

import (
	"context"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func (storage *PGstorage) AddUsers(ctx context.Context, infos []*models.UserInfo) error {
	query := storage.addUsersQuery(infos)
	queryText, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "generate !users! query error")
	}
	_, err = storage.db.Exec(ctx, queryText, args...)
	if err != nil {
		err = errors.Wrap(err, "exeс !users! query error")
	}
	return err
}

func (storage *PGstorage) addUsersQuery(userInfos []*models.UserInfo) squirrel.Sqlizer {
	infos := lo.Map(userInfos, func(info *models.UserInfo, _ int) *UserInfo {
		return &UserInfo{
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
	})

	q := squirrel.
		Insert(userTableName).
		Columns(
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
		PlaceholderFormat(squirrel.Dollar)
	for _, info := range infos {
		q = q.Values(
			info.Name,
			info.Email,
			info.Sex,
			info.Age,
			info.HeightCm,
			info.WeightKg,
			info.TargetWeightKg,
			info.TargetCalories,
			info.TargetProteins,
			info.TargetFats,
			info.TargetCarbs,
		)
	}
	return q
}
