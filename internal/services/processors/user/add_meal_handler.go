package userprocessor

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (p *processor) AddMealToUser(ctx context.Context, id uint64, mealInfo *models.MealInfo) error {
	return p.service.AddMealToUser(ctx, id, mealInfo)
}
