package meal

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (s *Service) GetByIds(ctx context.Context, ids []uint64) ([]*models.MealInfo, error) {
	return s.storage.GetMealsByIds(ctx, ids)
}
