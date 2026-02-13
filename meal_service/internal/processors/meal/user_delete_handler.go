package mealprocessor

import (
	"context"
)

func (p *Processor) CascadeDeletion(ctx context.Context, ids []uint64) error {
	return p.service.DeleteByUserIds(ctx, ids)
}
