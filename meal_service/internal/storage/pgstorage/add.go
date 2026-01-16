package pgstorage

import (
	"context"

	sharedmodels "github.com/JoePeach762/PP_project/shared/models"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func (storage *PGstorage) AddMeal(ctx context.Context, info *sharedmodels.MealInfo) error {
	query := storage.addMealsQuery([]*sharedmodels.MealInfo{info})
	queryText, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "generate !meals! single-query error")
	}
	_, err = storage.db.Exec(ctx, queryText, args...)
	if err != nil {
		err = errors.Wrap(err, "exeс !meals! single-query error")
	}
	return err
}

func (storage *PGstorage) AddMeals(ctx context.Context, infos []*sharedmodels.MealInfo) error {
	query := storage.addMealsQuery(infos)
	queryText, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "generate !meals! query error")
	}
	_, err = storage.db.Exec(ctx, queryText, args...)
	if err != nil {
		err = errors.Wrap(err, "exeс !meals! query error")
	}
	return err
}

func (storage *PGstorage) addMealsQuery(mealInfos []*sharedmodels.MealInfo) squirrel.Sqlizer {
	infos := lo.Map(mealInfos, func(info *sharedmodels.MealInfo, _ int) *MealInfo {
		return &MealInfo{
			ID:           info.ID,
			UserId:       info.UserId,
			Name:         info.Name,
			WeightGrams:  info.WeightGrams,
			Calories100g: info.Calories100g,
			Proteins100g: info.Proteins100g,
			Fats100g:     info.Fats100g,
			Carbs100g:    info.Carbs100g,
			Date:         info.Date,
		}
	})

	q := squirrel.
		Insert(mealTableName).
		Columns(
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
		PlaceholderFormat(squirrel.Dollar)
	for _, info := range infos {
		q = q.Values(
			info.ID,
			info.UserId,
			info.Name,
			info.WeightGrams,
			info.Calories100g,
			info.Proteins100g,
			info.Fats100g,
			info.Carbs100g,
			info.Date,
		)
	}
	return q
}
