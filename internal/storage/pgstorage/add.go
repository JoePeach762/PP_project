package pgstorage

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
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
			Name:            info.Name,
			Email:           info.Email,
			Sex:             info.Sex,
			Age:             info.Age,
			HeightCm:        info.HeightCm,
			WeightKg:        info.WeightKg,
			TargetWeightKg:  info.TargetWeightKg,
			CurrentCalories: 0,
			CurrentProteins: 0,
			CurrentFats:     0,
			CurrentCarbs:    0,
			TargetCalories:  info.TargetCalories,
			TargetProteins:  info.TargetProteins,
			TargetFats:      info.TargetFats,
			TargetCarbs:     info.TargetCarbs,
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
			userCurrentCaloriesColumnName,
			userCurrentProteinsColumnName,
			userCurrentFatsColumnName,
			userCurrentCarbsColumnName,
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
			info.CurrentCalories,
			info.CurrentProteins,
			info.CurrentFats,
			info.CurrentCarbs,
			info.TargetCalories,
			info.TargetProteins,
			info.TargetFats,
			info.TargetCarbs,
		)
	}
	return q
}

func (storage *PGstorage) AddMeal(ctx context.Context, info *models.MealInfo) error {
	query := storage.addMealsQuery([]*models.MealInfo{info})
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

func (storage *PGstorage) AddMeals(ctx context.Context, infos []*models.MealInfo) error {
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

func (storage *PGstorage) addMealsQuery(mealInfos []*models.MealInfo) squirrel.Sqlizer {
	infos := lo.Map(mealInfos, func(info *models.MealInfo, _ int) *MealInfo {
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
