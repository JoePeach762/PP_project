package mealproducer

import (
	"context"
	"encoding/json"

	"github.com/JoePeach762/PP_project/meal_service/internal/models"
)

func (p *MealKafkaProducer) PublishMealConsumed(ctx context.Context, meal *models.MealInfo) error {
	data, err := json.Marshal(meal)
	if err != nil {
		return err
	}

	return p.PublishMessage(ctx, data)
}
