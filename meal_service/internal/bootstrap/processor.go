package bootstrap

import (
	"context"

	mealprocessor "github.com/JoePeach762/PP_project/meal_service/internal/processors/meal"
	"github.com/JoePeach762/PP_project/meal_service/internal/services/meal"
)

func InitMealProcessor(mealService *meal.Service) *mealprocessor.Processor {
	return mealprocessor.NewMealProcessor(context.Background(), mealService)
}
