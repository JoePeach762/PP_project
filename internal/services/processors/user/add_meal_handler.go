package userprocessor

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (p *Processor) AddMealToUser(ctx context.Context, mealInfo *models.MealInfo) error {
	return p.service.AddMealToUser(ctx, mealInfo)
}
