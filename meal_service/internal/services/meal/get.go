package meal

import (
	"context"

	sharedmodels "github.com/JoePeach762/PP_project/shared/models"
)

func (s *Service) GetMealsByUserId(ctx context.Context, id uint64) ([]*sharedmodels.MealInfo, error) {
	return s.storage.GetMealsByUserId(ctx, id)
}
