package userprocessor

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (p *processor) Handle(ctx context.Context, info *models.UserInfo) error {
	return p.service.Add(ctx, []*models.UserInfo{info})
}
