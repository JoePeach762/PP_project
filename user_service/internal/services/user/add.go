package user

import (
	"context"
	"fmt"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
)

func (s *Service) Add(ctx context.Context, input []*models.UserInput) error {
	if err := s.Validate(input); err != nil {
		return err
	}
	infos := make([]*models.UserInfo, 0, len(input))
	for _, in := range input {
		infos = append(infos, models.NewUserInfoFromInput(in))
	}

	if err := s.calculateTargets(infos); err != nil {
		return fmt.Errorf("не удалось рассчитать цели: %w", err)
	}
	return s.userStorage.AddUsers(ctx, infos)
}
