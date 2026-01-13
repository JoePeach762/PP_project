package user

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (s *service) Update(ctx context.Context, id uint64, info models.UserInfo) error {
	if err := s.validateSingle(&info); err != nil {
		return err
	}
	return s.storage.UpdateUser(ctx, id, info)
}
