package user

import (
	"context"
	"fmt"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (s *service) Add(ctx context.Context, infos []*models.UserInfo) error {
	if err := s.Validate(infos); err != nil {
		return err
	}
	if err := s.calculateTargets(infos); err != nil {
		return fmt.Errorf("не удалось рассчитать цели: %w", err)
	}
	return s.storage.AddUsers(ctx, infos)
}
