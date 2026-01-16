package bootstrap

import (
	"github.com/JoePeach762/PP_project/meal_service/config"
	"github.com/JoePeach762/PP_project/meal_service/internal/services/meal"
)

func InitOFFClient(cfg *config.Config) meal.OFFClient {
	userAgent := cfg.MealServiceSettings.OFFUserAgent
	if userAgent == "" {
		userAgent = "PP_NutritionApp/1.0 (default@example.com)"
	}
	return meal.NewHTTPClient(userAgent)
}
