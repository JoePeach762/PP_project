package userprocessor

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

type service interface {
	AddMealToUser(ctx context.Context, info *models.MealInfo) error
}

type Processor struct {
	service service
}

func NewUserProcessor(ctx context.Context, service service) *Processor {
	return &Processor{
		service: service,
	}
}
