package userapi

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (a *UserAPI) Add(ctx context.Context, infos []*models.UserInfo) error {
	return a.service.Add(ctx, infos)
}
