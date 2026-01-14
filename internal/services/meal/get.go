package meal

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (s *service) GetOrCreateMealInfo(ctx context.Context, name string) (*models.MealInfo, error) {
	if meal, err := s.storage.GetMealByName(ctx, name); err == nil && meal != nil {
		return meal, nil
	}

	offProduct, err := s.offClient.FetchProduct(ctx, name)
	if err != nil {
		return nil, err
	}

	meal := &models.MealInfo{
		Name:         offProduct.Name,
		Calories100g: offProduct.Calories100g,
		Proteins100g: offProduct.Proteins100g,
		Fats100g:     offProduct.Fats100g,
		Carbs100g:    offProduct.Carbs100g,
	}

	s.storage.AddMeal(ctx, meal)

	return meal, nil
}
