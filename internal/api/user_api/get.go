package userapi

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (a *UserAPI) GetByIds(ctx context.Context, ids []uint64) ([]*models.UserInfo, error) {
	return a.service.GetByIds(ctx, ids)
}
