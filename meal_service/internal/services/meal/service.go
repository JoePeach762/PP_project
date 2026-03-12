package meal

import (
	"context"

	"github.com/JoePeach762/PP_project/meal_service/internal/models"
)

type OFFClient interface {
	FetchProduct(ctx context.Context, name string) (*models.MealTemplate, error)
}

type storage interface {
	AddMealAndEnqueueEvent(ctx context.Context, info *models.MealInfo) error
	GetMealsByUserId(ctx context.Context, id uint64) ([]*models.MealInfo, error)
	DeleteByUserIds(ctx context.Context, ids []uint64) error
}

type cache interface {
	AddProduct(ctx context.Context, info *models.MealTemplate) error
	GetProduct(ctx context.Context, name string) (*models.MealTemplate, error)
}

type Service struct {
	storage        storage
	cache          cache
	offClient      OFFClient
	minNameLength  uint32
	maxNameLength  uint32
	minWeightGrams uint32
	maxWeightGrams uint32
}

func NewMealService(
	storage storage,
	cache cache,
	offClient OFFClient,
	minNameLength uint32,
	maxNameLength uint32,
	minWeightGrams uint32,
	maxWeightGrams uint32,
) *Service {
	return &Service{
		storage:        storage,
		cache:          cache,
		offClient:      offClient,
		minNameLength:  minNameLength,
		maxNameLength:  maxNameLength,
		minWeightGrams: minWeightGrams,
		maxWeightGrams: maxWeightGrams,
	}
}
