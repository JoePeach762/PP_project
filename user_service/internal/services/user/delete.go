package user

import (
	"context"
)

func (s *Service) DeleteByIds(ctx context.Context, ids []uint64) error {
	return s.userStorage.DeleteUsersAndEnqueueEvent(ctx, ids)
}
