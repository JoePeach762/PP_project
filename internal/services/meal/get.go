package meal

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (s *service) Get(ctx context.Context, name []string) (*models.MealInfo, error) {
	return s.storage.GetMeal(ctx, name)
}
