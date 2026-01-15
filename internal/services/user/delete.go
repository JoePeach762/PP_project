package user

import (
	"context"
)

func (s *Service) DeleteUsers(ctx context.Context, ids []uint64) error {
	return s.storage.DeleteUsers(ctx, ids)
}
