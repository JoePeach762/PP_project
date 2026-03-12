package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/JoePeach762/PP_project/meal_service/internal/models"
)

func (c *RedisCache) AddProduct(ctx context.Context, info *models.MealTemplate) error {
	key := "product:" + info.Name
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, data, 72*time.Hour).Err()
}
