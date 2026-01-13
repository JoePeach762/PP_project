package user

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

type storage interface {
	AddUsers(ctx context.Context, infos []*models.UserInfo) error
	GetUsersByIds(ctx context.Context, ids []uint64) ([]*models.UserInfo, error)
	UpdateUser(ctx context.Context, id uint64, info models.UserInfo) error
	// DeleteUsers(ctx context.Context, ids []uint64) error
}

type service struct {
	storage       storage
	minNameLength uint8
	maxNameLength uint8
	minWeight     uint8
	maxWeight     uint8
}

func NewUserService(ctx context.Context,
	storage storage,
	minNameLength uint8,
	maxNameLength uint8,
	minWeight uint8,
	maxWeight uint8,
) *service {
	return &service{
		storage: storage,
	}
}
