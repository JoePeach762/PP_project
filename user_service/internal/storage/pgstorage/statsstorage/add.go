package statsstorage

import (
	"context"
	"fmt"
	"math"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s *PGstorage) AddMealToUser(ctx context.Context, info *models.MealInfo) (err error) {
	if info.EventID == "" {
		return errors.New("требуется event id события о приеме пищи")
	}

	stats := newStatsFromMeal(info)
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return errors.Wrap(err, "не удалось начать транзакцию добавления приема пищи")
	}
	defer func() {
		if err == nil {
			return
		}
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil && rollbackErr != pgx.ErrTxClosed {
			err = errors.Wrap(rollbackErr, "не удалось откатить транзакцию добавления приема пищи")
		}
	}()

	eventQuery := squirrel.Insert(processedMealEventsTableName).
		Columns(processedMealEventIDColumnName, processedMealUserIDColumnName).
		Values(info.EventID, info.UserId).
		Suffix(fmt.Sprintf("ON CONFLICT (%s) DO NOTHING", processedMealEventIDColumnName)).
		PlaceholderFormat(squirrel.Dollar)

	eventQueryText, eventArgs, err := eventQuery.ToSql()
	if err != nil {
		return errors.Wrap(err, "не удалось сформировать INSERT-запрос в processed_meal_events")
	}

	eventResult, err := tx.Exec(ctx, eventQueryText, eventArgs...)
	if err != nil {
		return errors.Wrap(err, "не удалось выполнить INSERT-запрос в processed_meal_events")
	}
	if eventResult.RowsAffected() == 0 {
		if err = tx.Commit(ctx); err != nil {
			return errors.Wrap(err, "не удалось зафиксировать транзакцию дубликата события о приеме пищи")
		}
		return nil
	}

	query := squirrel.Insert(statsTableName).
		Columns(
			statsUserIDColumnName,
			statsCaloriesColumnName,
			statsProteinsColumnName,
			statsFatsColumnName,
			statsCarbsColumnName,
			statsDateColumnName,
		).
		Values(
			stats.UserID,
			stats.CurrentCalories,
			stats.CurrentProteins,
			stats.CurrentFats,
			stats.CurrentCarbs,
			stats.Date,
		).
		Suffix(fmt.Sprintf(
			`ON CONFLICT (%s, %s) DO UPDATE SET
			%s = %s.%s + EXCLUDED.%s,
			%s = %s.%s + EXCLUDED.%s,
			%s = %s.%s + EXCLUDED.%s,
			%s = %s.%s + EXCLUDED.%s`,
			statsUserIDColumnName, statsDateColumnName,
			statsCaloriesColumnName, statsTableName, statsCaloriesColumnName, statsCaloriesColumnName,
			statsProteinsColumnName, statsTableName, statsProteinsColumnName, statsProteinsColumnName,
			statsFatsColumnName, statsTableName, statsFatsColumnName, statsFatsColumnName,
			statsCarbsColumnName, statsTableName, statsCarbsColumnName, statsCarbsColumnName,
		)).
		PlaceholderFormat(squirrel.Dollar)

	queryText, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "не удалось сформировать INSERT-запрос добавления приема пищи")
	}

	_, err = tx.Exec(ctx, queryText, args...)
	if err != nil {
		return errors.Wrap(err, "не удалось выполнить INSERT-запрос добавления приема пищи")
	}

	if err = tx.Commit(ctx); err != nil {
		return errors.Wrap(err, "не удалось зафиксировать транзакцию добавления приема пищи")
	}

	return nil
}

func newStatsFromMeal(info *models.MealInfo) *models.UserStats {
	return &models.UserStats{
		UserID:          info.UserId,
		CurrentCalories: uint32(math.Round(float64(info.Calories100g * info.WeightGrams / 100))),
		CurrentProteins: uint32(math.Round(float64(info.Proteins100g * info.WeightGrams / 100))),
		CurrentFats:     uint32(math.Round(float64(info.Fats100g * info.WeightGrams / 100))),
		CurrentCarbs:    uint32(math.Round(float64(info.Carbs100g * info.WeightGrams / 100))),
		Date:            normalizeDate(info.Date),
	}
}
