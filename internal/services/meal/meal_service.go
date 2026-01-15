package meal

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

type OFFClient interface {
	FetchProduct(ctx context.Context, name string) (*models.MealTemplate, error)
}

type producer interface {
	PublishMealConsumed(ctx context.Context, event *models.MealInfo) error
}

type storage interface {
	AddMeal(ctx context.Context, info *models.MealInfo) error
	GetMealsByIds(ctx context.Context, ids []uint64) ([]*models.MealInfo, error)
}

type cache interface {
	AddProduct(ctx context.Context, info *models.MealTemplate) error
	GetProduct(ctx context.Context, name string) (*models.MealTemplate, error)
}

type Service struct {
	producer       producer
	storage        storage
	cache          cache
	offClient      OFFClient
	minNameLength  uint32
	maxNameLength  uint32
	maxWeightGrams uint32
}

func NewMealService(
	ctx context.Context,
	producer producer,
	storage storage,
	cache cache,
	offClient OFFClient,
	minNameLength uint32,
	maxNameLength uint32,
	maxWeightGrams uint32,
) *Service {
	return &Service{
		producer:       producer,
		storage:        storage,
		cache:          cache,
		offClient:      offClient,
		minNameLength:  minNameLength,
		maxNameLength:  maxNameLength,
		maxWeightGrams: maxWeightGrams,
	}
}
