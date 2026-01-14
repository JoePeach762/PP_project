package pgstorage

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (storage *PGstorage) GetStudentInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.StudentInfo, error) {
	query := storage.getQuery(IDs)
	queryText, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "generate query error")
	}
	rows, err := storage.db.Query(ctx, queryText, args...)
	if err != nil {
		return nil, errors.Wrap(err, "quering error")
	}
	var students []*models.StudentInfo
	for rows.Next() {
		var s models.StudentInfo
		if err := rows.Scan(&s.ID, &s.Name, &s.Email, &s.Age); err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}
		students = append(students, &s)
	}
	return students, nil
}

func (storage *PGstorage) getQuery(IDs []uint64) squirrel.Sqlizer {
	q := squirrel.Select(studentIDColumnName, studentNameColumnName, studentEmailColumnName, studentAgeColumnName).From(studentTableName).
		Where(squirrel.Eq{studentIDColumnName: IDs}).PlaceholderFormat(squirrel.Dollar)
	return q
}

func (storage *PGstorage) GetUsersByIds(ctx context.Context, ids []uint64) ([]*models.UserInfo, error) {
	query := storage.getUsersQuery(ids)
	queryText, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "generate !users! query error")
	}
	rows, err := storage.db.Query(ctx, queryText, args...)
	if err != nil {
		return nil, errors.Wrap(err, "quering !users! error")
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
			&u.CurrentCalories,
			&u.CurrentProteins,
			&u.CurrentFats,
			&u.CurrentCarbs,
			&u.TargetCalories,
			&u.TargetProteins,
			&u.TargetFats,
			&u.TargetCarbs,
		); err != nil {
			return nil, errors.Wrap(err, "failed to scan !users! row")
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
			userCurrentCaloriesColumnName,
			userCurrentProteinsColumnName,
			userCurrentFatsColumnName,
			userCurrentCarbsColumnName,
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
