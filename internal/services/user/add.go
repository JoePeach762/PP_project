package user

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (s *service) Add(ctx context.Context, infos []*models.UserInfo) error {
	if err := s.Validate(infos); err != nil {
		return err
	}
	return s.storage.AddUsers(ctx, infos)
}
