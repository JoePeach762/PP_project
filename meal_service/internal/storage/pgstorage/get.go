package pgstorage

import (
	"context"

	sharedmodels "github.com/JoePeach762/PP_project/shared/models"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (storage *PGstorage) GetMealsByUserId(ctx context.Context, id uint64) ([]*sharedmodels.MealInfo, error) {
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

	var meals []*sharedmodels.MealInfo
	for rows.Next() {
		var m sharedmodels.MealInfo
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
