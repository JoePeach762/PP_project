package userapi

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

type service interface {
	Add(ctx context.Context, infos []*models.UserInfo) error
	GetByIds(ctx context.Context, ids []uint64) ([]*models.UserInfo, error)
	Update(ctx context.Context, id uint64, info models.UserInfo) error
	DeleteByIds(ctx context.Context, ids []uint64) error
}

type UserAPI struct {
	service service
}

func NewUserAPI(ctx context.Context, service service) *UserAPI {
	return &UserAPI{
		service: service,
	}
}
