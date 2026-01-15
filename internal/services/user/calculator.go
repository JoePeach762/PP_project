package user

import (
	"errors"
	"math"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (s *Service) calculateTargets(infos []*models.UserInfo) error {
	for _, info := range infos {
		if err := s.calculateTargetsSingle(info); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) calculateTargetsSingle(info *models.UserInfo) error {
	if info.WeightKg == 0 || info.HeightCm == 0 || info.Age == 0 {
		return errors.New("недостаточно данных для расчёта целей")
	}

	bmr := 10*float64(info.WeightKg) + 6.25*float64(info.HeightCm) - 5*float64(info.Age)
	if info.Sex == "female" {
		bmr -= 161
	} else {
		bmr += 5
	}

	tdee := bmr * 1.375

	var targetCalories float64
	if info.TargetWeightKg < info.WeightKg {
		targetCalories = tdee - 500
	} else if info.TargetWeightKg > info.WeightKg {
		targetCalories = tdee + 300
	} else {
		targetCalories = tdee
	}

	if targetCalories < 1200 {
		targetCalories = 1200
	}

	targetWeight := float64(info.TargetWeightKg)
	proteinsG := 2.0 * targetWeight
	fatsKcal := targetCalories * 0.275
	fatsG := fatsKcal / 9
	carbsKcal := targetCalories - proteinsG*4 - fatsG*9
	if carbsKcal < 0 {
		carbsKcal = 0
	}
	carbsG := carbsKcal / 4

	info.TargetCalories = uint32(math.Round(targetCalories))
	info.TargetProteins = uint32(math.Round(proteinsG))
	info.TargetFats = uint32(math.Round(fatsG))
	info.TargetCarbs = uint32(math.Round(carbsG))

	return nil
}
