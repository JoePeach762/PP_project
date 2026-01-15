package mealapi

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (a *MealAPI) Add(ctx context.Context, req *models.MealInput) error {
	return a.service.Add(ctx, req)
}
