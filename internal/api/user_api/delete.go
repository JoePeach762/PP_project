package userapi

import (
	"context"
)

func (a *UserAPI) DeleteByIds(ctx context.Context, ids []uint64) error {
	return a.service.DeleteByIds(ctx, ids)
}
