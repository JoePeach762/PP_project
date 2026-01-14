package user

import (
	"context"
)

func (s *service) DeleteUsers(ctx context.Context, ids []uint64) error {
	return s.storage.DeleteUsers(ctx, ids)
}
