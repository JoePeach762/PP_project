package pgstorage

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

func (storage *PGstorage) GetMealsByUserId(ctx context.Context, id uint64) ([]*models.MealInfo, error) {
	query := storage.getMealsQuery(id)
	queryText, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "generate !meals! query error")
	}
	rows, err := storage.db.Query(ctx, queryText, args...)
	if err != nil {
		return nil, errors.Wrap(err, "quering !meals! error")
	}
	defer rows.Close()

	var meals []*models.MealInfo
	for rows.Next() {
		var m models.MealInfo
		if err := rows.Scan(
			&m.ID,
			&m.UserId,
			&m.Name,
			&m.WeightGrams,
			&m.Calories100g,
			&m.Proteins100g,
			&m.Fats100g,
			&m.Carbs100g,
			&m.Date,
		); err != nil {
			return nil, errors.Wrap(err, "failed to scan !meals! row")
		}
		meals = append(meals, &m)
	}
	return meals, nil
}

func (storage *PGstorage) getMealsQuery(UserId uint64) squirrel.Sqlizer {
	q := squirrel.
		Select(
			mealIDColumnName,
			mealUserIDcolumnName,
			mealNameColumnName,
			mealWeightGramsColumnName,
			mealCalories100gColumnName,
			mealProteins100gColumnName,
			mealFats100gColumnName,
			mealCarbs100gColumnName,
			mealDateColumnName,
		).
		From(mealTableName).
		Where(squirrel.Eq{mealUserIDcolumnName: UserId}).
		PlaceholderFormat(squirrel.Dollar)
	return q
}
