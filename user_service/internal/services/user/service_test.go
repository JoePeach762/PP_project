package user

import (
	"testing"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
)

func newServiceForTest(userStorage userStorage, statsStorage statsStorage) *Service {
	return NewUserService(
		userStorage,
		statsStorage,
		2,
		30,
		1,
		120,
		50,
		250,
		20,
		300,
		[]string{" male ", "FEMALE"},
	)
}

func newPermissiveServiceForTest(userStorage userStorage, statsStorage statsStorage) *Service {
	return NewUserService(
		userStorage,
		statsStorage,
		1,
		30,
		0,
		120,
		0,
		250,
		0,
		300,
		[]string{"male", "female"},
	)
}

func validUserInput() *models.UserInput {
	return &models.UserInput{
		Name:           "Alice",
		Email:          "alice@example.com",
		Sex:            "Female",
		Age:            28,
		HeightCm:       165,
		WeightKg:       60,
		TargetWeightKg: 55,
	}
}

func TestNewUserService_NormalizesAllowedSexes(t *testing.T) {
	service := NewUserService(nil, nil, 2, 30, 1, 120, 50, 250, 20, 300, []string{" Male ", "FEMALE"})

	if len(service.allowedSexes) != 2 {
		t.Fatalf("неожиданное количество допустимых значений пола: %d", len(service.allowedSexes))
	}

	if _, ok := service.allowedSexes["male"]; !ok {
		t.Fatalf("ожидалось, что male будет нормализован в allowedSexes")
	}

	if _, ok := service.allowedSexes["female"]; !ok {
		t.Fatalf("ожидалось, что female будет нормализован в allowedSexes")
	}
}
