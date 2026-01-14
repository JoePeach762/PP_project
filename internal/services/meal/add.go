package meal

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (s *service) Add(ctx context.Context, info *models.MealInfo) error {
	return s.storage.AddMeal(ctx, info)
}
