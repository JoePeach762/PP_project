package mealapi

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (a *MealAPI) GetByIds(ctx context.Context, ids []uint64) ([]*models.MealInfo, error) {
	return a.service.GetByIds(ctx, ids)
}
