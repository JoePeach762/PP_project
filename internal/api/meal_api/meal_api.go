package mealapi

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

type service interface {
	Add(ctx context.Context, req *models.MealInput) error
	GetByIds(ctx context.Context, ids []uint64) ([]*models.MealInfo, error)
}

type MealAPI struct {
	service service
}

func NewMealAPI(ctx context.Context, service service) *MealAPI {
	return &MealAPI{
		service: service,
	}
}
