package userprocessor

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

type service interface {
	Add(ctx context.Context, infos []*models.UserInfo) error
	GetAll(ctx context.Context) ([]*models.UserInfo, error)
	GetById(ctx context.Context, ids []uint64) ([]*models.UserInfo, error)
	Update(ctx context.Context, id uint64, info models.UserInfo) error
	Delete(ctx context.Context, id uint64) error
}

type processor struct {
	service service
}

func NewUserProcessor(ctx context.Context, service service) *processor {
	return &processor{
		service: service,
	}
}
