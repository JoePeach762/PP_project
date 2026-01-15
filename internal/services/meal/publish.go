package meal

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (s *Service) Publish(ctx context.Context, event *models.MealInfo) error {
	return s.producer.PublishMealConsumed(ctx, event)
}
