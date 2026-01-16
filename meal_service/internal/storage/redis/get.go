package redis

import (
	"context"
	"encoding/json"

	"github.com/JoePeach762/PP_project/meal_service/internal/models"
)

func (c *RedisCache) GetProduct(ctx context.Context, name string) (*models.MealTemplate, error) {
	key := "product:" + name
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var product models.MealTemplate
	err = json.Unmarshal([]byte(val), &product)
	return &product, err
}
