package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (c *RedisCache) AddProduct(ctx context.Context, info *models.MealTemplate) error {
	key := "product:" + info.Name
	data, _ := json.Marshal(info)
	return c.client.Set(ctx, key, data, 72*time.Hour).Err()
}
