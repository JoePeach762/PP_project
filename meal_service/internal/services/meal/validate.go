package meal

import (
	"fmt"

	"github.com/JoePeach762/PP_project/meal_service/internal/models"
)

func (s *Service) validate(req *models.MealInput) error {
	if len(req.Name) < int(s.minNameLength) || len(req.Name) > int(s.maxNameLength) {
		return fmt.Errorf("имя должно быть длиной от %d до %d символов", s.minNameLength, s.maxNameLength)
	}
	if req.WeightGrams < float32(s.minWeightGrams) || req.WeightGrams > float32(s.maxWeightGrams) {
		return fmt.Errorf("вес должен быть в диапазоне от %d до %d граммов", s.minWeightGrams, s.maxWeightGrams)
	}
	return nil
}
