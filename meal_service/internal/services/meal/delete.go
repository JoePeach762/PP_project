package meal

import (
	"context"
)

func (s *Service) DeleteByUserIds(ctx context.Context, ids []uint64) error {
	return s.storage.DeleteByUserIds(ctx, ids)
}
