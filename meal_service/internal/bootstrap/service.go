package bootstrap

import (
	"github.com/JoePeach762/PP_project/meal_service/config"
	"github.com/JoePeach762/PP_project/meal_service/internal/services/meal"
	mealstorage "github.com/JoePeach762/PP_project/meal_service/internal/storage/pgstorage/mealstorage"
	redisstore "github.com/JoePeach762/PP_project/meal_service/internal/storage/redis"
)

func InitMealService(
	storage *mealstorage.PGstorage,
	cache *redisstore.RedisCache,
	offClient meal.OFFClient,
	cfg *config.Config,
) *meal.Service {
	return meal.NewMealService(
		storage,
		cache,
		offClient,
		cfg.MealServiceSettings.MinNameLen,
		cfg.MealServiceSettings.MaxNameLen,
		cfg.MealServiceSettings.MinWeightGrams,
		cfg.MealServiceSettings.MaxWeightGrams,
	)
}
