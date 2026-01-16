package meal

import (
	"errors"
	"fmt"

	"github.com/JoePeach762/PP_project/meal_service/internal/models"
)

func (s *Service) validate(req *models.MealInput) error {
	if len(req.Name) <= int(s.minNameLength) || len(req.Name) >= int(s.maxNameLength) {
		return errors.New("имя не должно быть пустым и не должно превышать 100 символов")
	}
	if req.WeightGrams <= float32(0) || req.WeightGrams >= float32(s.maxWeightGrams) {
		return fmt.Errorf("некорректный вес %v", req.WeightGrams)
	}
	return nil
}
