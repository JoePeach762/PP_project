package user

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (s *service) validateSingle(info *models.UserInfo) error {
	if len(info.Name) <= int(s.minNameLength) || len(info.Name) >= int(s.maxNameLength) {
		return errors.New("имя не должно быть пустым и не должно превышать 100 символов")
	}
	if info.Age <= 0 || info.Age > 100 {
		return fmt.Errorf("некорректный возраст %v", info.Age)
	}
	if info.HeightCm <= 0 || info.HeightCm > 300 {
		return fmt.Errorf("некорректный рост %v", info.HeightCm)
	}
	if info.WeightKg <= 0 || info.WeightKg > 200 {
		return fmt.Errorf("некорректный вес %v", info.WeightKg)
	}
	if info.TargetWeightKg <= 0 || info.TargetWeightKg > 200 {
		return fmt.Errorf("некорректный целевой вес %v", info.TargetWeightKg)
	}
	if strings.ToLower(info.Sex) != "male" && strings.ToLower(info.Sex) != "female" {
		return fmt.Errorf("некорректный пол %v", info.Sex)
	}
	if !s.isValidEmail(info.Email) {
		return fmt.Errorf("некорректный email: %v", info.Email)
	}
	return nil
}

func (s *service) Validate(infos []*models.UserInfo) error {
	for _, info := range infos {
		if err := s.validateSingle(info); err != nil {
			return err
		}
	}
	return nil
}

func (s *service) isValidEmail(email string) bool {
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
