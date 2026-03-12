package user

import (
	"fmt"
	"net/mail"
	"strings"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
)

func (s *Service) validateSingle(info *models.UserInput) error {
	if len(info.Name) < int(s.minNameLength) || len(info.Name) > int(s.maxNameLength) {
		return fmt.Errorf("имя должно быть длиной от %d до %d символов", s.minNameLength, s.maxNameLength)
	}
	if info.Age < s.minAge || info.Age > s.maxAge {
		return fmt.Errorf("возраст должен быть в диапазоне от %d до %d", s.minAge, s.maxAge)
	}
	if info.HeightCm < s.minHeightCm || info.HeightCm > s.maxHeightCm {
		return fmt.Errorf("рост должен быть в диапазоне от %d до %d см", s.minHeightCm, s.maxHeightCm)
	}
	if info.WeightKg < s.minWeight || info.WeightKg > s.maxWeight {
		return fmt.Errorf("вес должен быть в диапазоне от %d до %d кг", s.minWeight, s.maxWeight)
	}
	if info.TargetWeightKg < s.minWeight || info.TargetWeightKg > s.maxWeight {
		return fmt.Errorf("целевой вес должен быть в диапазоне от %d до %d кг", s.minWeight, s.maxWeight)
	}
	normalizedSex := strings.ToLower(strings.TrimSpace(info.Sex))
	if _, ok := s.allowedSexes[normalizedSex]; !ok {
		return fmt.Errorf("некорректный пол %v", info.Sex)
	}
	info.Sex = normalizedSex
	if !s.isValidEmail(info.Email) {
		return fmt.Errorf("некорректный email: %v", info.Email)
	}
	return nil
}

func (s *Service) Validate(infos []*models.UserInput) error {
	for _, info := range infos {
		if err := s.validateSingle(info); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) isValidEmail(email string) bool {
	if len(email) < 3 || len(email) > 254 {
		return false
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	if len(parts[1]) == 0 || len(parts[1]) > 253 {
		return false
	}

	return true
}
