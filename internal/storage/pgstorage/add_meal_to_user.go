package pgstorage

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (s *PGstorage) AddMealToUser(ctx context.Context, info *models.MealInfo) error {
	query := squirrel.Update(userTableName).
		Set(userCurrentCaloriesColumnName, info.Calories100g*info.WeightGrams/100).
		Set(userCurrentProteinsColumnName, info.Proteins100g*info.WeightGrams/100).
		Set(userCurrentFatsColumnName, info.Fats100g*info.WeightGrams/100).
		Set(userCurrentCarbsColumnName, info.Carbs100g*info.WeightGrams/100).
		Where(squirrel.Eq{userIDColumnName: info.UserId}).
		PlaceholderFormat(squirrel.Dollar)

	queryText, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "generate addMeal UPDATE query")
	}

	result, err := s.db.Exec(ctx, queryText, args...)
	if err != nil {
		return errors.Wrap(err, "execute addMeal UPDATE query")
	}

	if result.RowsAffected() == 0 {
		return errors.New("addMeal failed: user not found")
	}

	return nil
}
