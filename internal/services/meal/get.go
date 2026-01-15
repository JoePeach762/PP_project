package meal

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (s *Service) GetMealsByUserId(ctx context.Context, id uint64) ([]*models.MealInfo, error) {
	return s.storage.GetMealsByUserId(ctx, id)
}
