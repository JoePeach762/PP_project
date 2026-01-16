package meal

import (
	"context"
	"time"

	"github.com/JoePeach762/PP_project/meal_service/internal/models"
)

func (s *Service) Add(ctx context.Context, req *models.MealInput) error {
	if err := s.validate(req); err != nil {
		return err
	}

	template, err := s.cache.GetProduct(ctx, req.Name)
	if err != nil {
		template, err = s.offClient.FetchProduct(ctx, req.Name)
		if err != nil {
			return err
		}
		s.cache.AddProduct(ctx, template)
	}

	info := &models.MealInfo{
		Name:         req.Name,
		UserId:       req.UserID,
		WeightGrams:  req.WeightGrams,
		Calories100g: template.Calories100g,
		Proteins100g: template.Proteins100g,
		Fats100g:     template.Fats100g,
		Carbs100g:    template.Carbs100g,
		Date:         time.Now(),
	}

	err = s.storage.AddMeal(ctx, info)
	if err != nil {
		return err
	}
	return s.producer.PublishMealConsumed(ctx, info)
}
