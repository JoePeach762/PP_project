package user

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

type storage interface {
	AddUsers(ctx context.Context, infos []*models.UserInfo) error
	GetUsersByIds(ctx context.Context, ids []uint64) ([]*models.UserInfo, error)
	UpdateUser(ctx context.Context, id uint64, info models.UserInfo) error
	DeleteUsers(ctx context.Context, ids []uint64) error
	AddMealToUser(ctx context.Context, id uint64, mealInfo models.MealInfo) error
}

type service struct {
	storage       storage
	minNameLength uint32
	maxNameLength uint32
	minWeight     uint32
	maxWeight     uint32
}

func NewUserService(ctx context.Context,
	storage storage,
	minNameLength uint32,
	maxNameLength uint32,
	minWeight uint32,
	maxWeight uint32,
) *service {
	return &service{
		storage:       storage,
		minNameLength: minNameLength,
		maxNameLength: maxNameLength,
		minWeight:     minWeight,
		maxWeight:     maxWeight,
	}
}
