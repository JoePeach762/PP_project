package bootstrap

import (
	"github.com/JoePeach762/PP_project/config"
	"github.com/JoePeach762/PP_project/internal/services/meal"
)

func InitOFFClient(cfg *config.Config) meal.OFFClient {
	userAgent := cfg.MealServiceSettings.OFFUserAgent
	if userAgent == "" {
		userAgent = "PP_NutritionApp/1.0 (default@example.com)"
	}
	return meal.NewHTTPClient(userAgent)
}
