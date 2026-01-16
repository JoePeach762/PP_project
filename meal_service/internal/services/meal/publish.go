package meal

import (
	"context"

	sharedmodels "github.com/JoePeach762/PP_project/shared/models"
)

func (s *Service) Publish(ctx context.Context, event *sharedmodels.MealInfo) error {
	return s.producer.PublishMealConsumed(ctx, event)
}
