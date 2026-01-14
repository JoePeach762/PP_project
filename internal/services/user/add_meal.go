package user

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (s *service) AddMealToUser(ctx context.Context, id uint64, meal models.MealInfo) error {
	return s.storage.AddMealToUser(ctx, id, meal)
}
