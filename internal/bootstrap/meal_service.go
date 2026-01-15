package bootstrap

import (
	"github.com/JoePeach762/PP_project/config"
	mealproducer "github.com/JoePeach762/PP_project/internal/producer/meal"
	"github.com/JoePeach762/PP_project/internal/services/meal"
	"github.com/JoePeach762/PP_project/internal/storage/pgstorage"
	redisstore "github.com/JoePeach762/PP_project/internal/storage/redis"
)

func InitMealService(
	storage *pgstorage.PGstorage,
	cache *redisstore.RedisCache,
	producer *mealproducer.MealKafkaProducer,
	offClient meal.OFFClient,
	cfg *config.Config,
) *meal.Service {
	return meal.NewMealService(
		producer,
		storage,
		cache,
		offClient,
		cfg.MealServiceSettings.MinNameLen,
		cfg.MealServiceSettings.MaxNameLen,
		cfg.MealServiceSettings.MaxWeightGrams,
	)
}
