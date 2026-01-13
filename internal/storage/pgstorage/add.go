package pgstorage

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func (storage *PGstorage) UpsertStudentInfo(ctx context.Context, studentInfos []*models.StudentInfo) error {
	query := storage.upsertQuery(studentInfos)
	queryText, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "generate query error")
	}
	_, err = storage.db.Exec(ctx, queryText, args...)
	if err != nil {
		err = errors.Wrap(err, "exeс query")
	}
	return err
}

func (storage *PGstorage) upsertQuery(studentInfos []*models.StudentInfo) squirrel.Sqlizer {
	infos := lo.Map(studentInfos, func(info *models.StudentInfo, _ int) *StudentInfo {
		return &StudentInfo{
			Name:  info.Name,
			Email: info.Email,
			Age:   info.Age,
		}
	})

	q := squirrel.Insert(studentTableName).Columns(studentNameColumnName, studentEmailColumnName, studentAgeColumnName).
		PlaceholderFormat(squirrel.Dollar)
	for _, info := range infos {
		q = q.Values(info.Name, info.Email, info.Age)
	}
	return q
}

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
			Name:            info.Name,
			Email:           info.Email,
			Sex:             info.Sex,
			Age:             info.Age,
			WeightKg:        info.WeightKg,
			TargetWeightKg:  info.TargetWeightKg,
			CurrentCalories: 0,
			CurrentProteins: 0,
			CurrentFats:     0,
			CurrentCarbs:    0,
		}
	})

	q := squirrel.
		Insert(userTableName).
		Columns(
			userNameColumnName,
			userEmailColumnName,
			userSexColumnName,
			userAgeColumnName,
			userWeightKgColumnName,
			userTargetWeightKgColumnName,
			userCurrentCaloriesColumnName,
			userCurrentProteinsColumnName,
			userCurrentFatsColumnName,
			userCurrentCarbsColumnName,
		).
		PlaceholderFormat(squirrel.Dollar)
	for _, info := range infos {
		q = q.Values(
			info.Name,
			info.Email,
			info.Sex,
			info.Age,
			info.WeightKg,
			info.TargetWeightKg,
			info.CurrentCalories,
			info.CurrentProteins,
			info.CurrentFats,
			info.CurrentCarbs,
		)
	}
	return q
}
