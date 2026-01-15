package user

import (
	"context"
	"fmt"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (s *Service) Update(ctx context.Context, id uint64, info models.UserInfo) error {
	if err := s.validateSingle(&info); err != nil {
		return err
	}
	if err := s.calculateTargetsSingle(&info); err != nil {
		return fmt.Errorf("не удалось пересчитать цели: %w", err)
	}
	return s.storage.UpdateUser(ctx, id, info)
}
