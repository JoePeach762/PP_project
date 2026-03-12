package mealstorage

import (
	"context"
	"encoding/json"

	"github.com/JoePeach762/PP_project/meal_service/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

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

func (storage *PGstorage) AddMealAndEnqueueEvent(ctx context.Context, info *models.MealInfo) (err error) {
	payload, err := json.Marshal(info)
	if err != nil {
		return errors.Wrap(err, "marshal meal event")
	}

	tx, err := storage.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return errors.Wrap(err, "begin add meal transaction")
	}
	defer func() {
		if err == nil {
			return
		}
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil && rollbackErr != pgx.ErrTxClosed {
			err = errors.Wrap(rollbackErr, "rollback add meal transaction")
		}
	}()

	query := storage.addMealsQuery([]*models.MealInfo{info})
	queryText, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "generate add meal query")
	}
	if _, err = tx.Exec(ctx, queryText, args...); err != nil {
		return errors.Wrap(err, "execute add meal query")
	}

	outboxQuery := squirrel.Insert(outboxTableName).
		Columns(outboxPayloadColumnName).
		Values(payload).
		PlaceholderFormat(squirrel.Dollar)

	outboxQueryText, outboxArgs, err := outboxQuery.ToSql()
	if err != nil {
		return errors.Wrap(err, "generate meal outbox query")
	}
	if _, err = tx.Exec(ctx, outboxQueryText, outboxArgs...); err != nil {
		return errors.Wrap(err, "execute meal outbox query")
	}

	if err = tx.Commit(ctx); err != nil {
		return errors.Wrap(err, "commit add meal transaction")
	}

	return nil
}

func (storage *PGstorage) addMealsQuery(mealInfos []*models.MealInfo) squirrel.Sqlizer {
	infos := lo.Map(mealInfos, func(info *models.MealInfo, _ int) *MealInfo {
		return &MealInfo{
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
