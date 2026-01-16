package meal

import (
	"context"

	"github.com/JoePeach762/PP_project/meal_service/internal/models"
)

func (s *Service) GetMealsByUserId(ctx context.Context, id uint64) ([]*models.MealInfo, error) {
	return s.storage.GetMealsByUserId(ctx, id)
}
