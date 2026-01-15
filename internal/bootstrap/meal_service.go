package bootstrap

import (
	"context"

	"github.com/JoePeach762/PP_project/config"
	mealproducer "github.com/JoePeach762/PP_project/internal/producer/meal"
	"github.com/JoePeach762/PP_project/internal/services/meal"
	"github.com/JoePeach762/PP_project/internal/storage/pgstorage"
	"github.com/JoePeach762/PP_project/internal/storage/redis"
)

func InitMealService(
	storage *pgstorage.PGstorage,
	cache *redis.RedisCache,
	producer *mealproducer.MealKafkaProducer,
	offClient meal.OFFClient,
	cfg *config.Config,
) *meal.Service {
	return meal.NewMealService(
		context.Background(),
		producer,
		storage,
		cache,
		offClient,
		cfg.MealServiceSettings.MinNameLen,
		cfg.MealServiceSettings.MaxNameLen,
		cfg.MealServiceSettings.MaxWeightGrams,
	)
}
