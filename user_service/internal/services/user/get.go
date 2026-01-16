package user

import (
	"context"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
)

func (s *Service) GetByIds(ctx context.Context, ids []uint64) ([]*models.UserInfo, error) {
	return s.storage.GetUsersByIds(ctx, ids)
}
