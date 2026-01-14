package meal

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

type producer interface {
	Publish() error
}

type storage interface {
	AddMeal(ctx context.Context, info *models.MealInfo) error
	GetMeal(ctx context.Context, name []string) (*models.MealInfo, error)
}

type service struct {
	producer       producer
	storage        storage
	minNameLength  uint32
	maxNameLength  uint32
	maxWeightGrams uint32
}

func NewMealService(
	ctx context.Context,
	producer producer,
	storage storage,
	minNameLength uint32,
	maxNameLength uint32,
	maxWeightGrams uint32,
) *service {
	return &service{
		producer:       producer,
		storage:        storage,
		minNameLength:  minNameLength,
		maxNameLength:  maxNameLength,
		maxWeightGrams: maxWeightGrams,
	}
}
