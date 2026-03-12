package user

import (
	"context"
	"fmt"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
)

func (s *Service) Update(ctx context.Context, id uint64, input models.UserInput) error {
	if err := s.validateSingle(&input); err != nil {
		return err
	}
	info := models.NewUserInfoFromInput(&input)
	if err := s.calculateTargetsSingle(info); err != nil {
		return fmt.Errorf("не удалось пересчитать цели: %w", err)
	}
	return s.userStorage.UpdateUser(ctx, id, *info)
}
