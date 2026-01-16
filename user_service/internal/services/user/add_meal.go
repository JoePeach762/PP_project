package user

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (s *Service) AddMealToUser(ctx context.Context, meal *models.MealInfo) error {
	return s.storage.AddMealToUser(ctx, meal)
}
