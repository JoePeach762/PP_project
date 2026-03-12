package user

import (
	"context"
)

func (s *Service) DeleteByIds(ctx context.Context, ids []uint64) error {
	err := s.userStorage.DeleteUsers(ctx, ids)
	if err != nil {
		return err
	}
	return s.producer.PublishUsersDeleted(ctx, ids)
}
