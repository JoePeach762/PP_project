package userapi

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (a *UserAPI) Update(ctx context.Context, id uint64, info models.UserInfo) error {
	return a.service.Update(ctx, id, info)
}
