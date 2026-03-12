package mealprocessor

import (
	"context"
)

type service interface {
	DeleteByUserIds(ctx context.Context, ids []uint64) error
}

type Processor struct {
	service service
}

func NewMealProcessor(ctx context.Context, service service) *Processor {
	return &Processor{
		service: service,
	}
}
