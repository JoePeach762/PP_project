package userprocessor

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

type service interface {
	AddMealToUser(ctx context.Context, id uint64, meal *models.MealInfo) error
}

type processor struct {
	service service
}

func NewUserProcessor(ctx context.Context, service service) *processor {
	return &processor{
		service: service,
	}
}
