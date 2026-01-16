package user

import (
	"context"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
)

type storage interface {
	AddUsers(ctx context.Context, infos []*models.UserInfo) error
	GetUsersByIds(ctx context.Context, ids []uint64) ([]*models.UserInfo, error)
	UpdateUser(ctx context.Context, id uint64, info models.UserInfo) error
	DeleteUsers(ctx context.Context, ids []uint64) error
	AddMealToUser(ctx context.Context, mealInfo *models.MealInfo) error
}

type Service struct {
	storage       storage
	minNameLength uint32
	maxNameLength uint32
	minWeight     uint32
	maxWeight     uint32
}

func NewUserService(
	storage storage,
	minNameLength uint32,
	maxNameLength uint32,
	minWeight uint32,
	maxWeight uint32,
) *Service {
	return &Service{
		storage:       storage,
		minNameLength: minNameLength,
		maxNameLength: maxNameLength,
		minWeight:     minWeight,
		maxWeight:     maxWeight,
	}
}
