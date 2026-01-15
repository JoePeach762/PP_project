package bootstrap

import (
	"github.com/JoePeach762/PP_project/internal/services/meal"
)

func InitOFFClient(userAgent string) meal.OFFClient {
	return meal.NewHTTPClient(userAgent)
}
