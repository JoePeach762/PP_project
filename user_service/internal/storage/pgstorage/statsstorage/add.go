package statsstorage

import (
	"context"
	"fmt"
	"math"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (s *PGstorage) AddMealToUser(ctx context.Context, info *models.MealInfo) error {
	stats := newStatsFromMeal(info)

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
		return errors.Wrap(err, "generate addMeal INSERT query")
	}

	_, err = s.db.Exec(ctx, queryText, args...)
	if err != nil {
		return errors.Wrap(err, "execute addMeal INSERT query")
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
