package user

import (
	"context"
	"strings"
	"testing"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
	"github.com/JoePeach762/PP_project/user_service/internal/services/user/mocks"
	"github.com/stretchr/testify/mock"
)

func TestServiceUpdate_SavesCalculatedUser(t *testing.T) {
	ctx := context.Background()
	userStorage := &mocks.UserStorage{}
	service := newServiceForTest(userStorage, nil)
	input := *validUserInput()

	userStorage.EXPECT().
		UpdateUser(ctx, uint64(7), mock.Anything).
		Run(func(_ context.Context, id uint64, info models.UserInfo) {
			if id != 7 {
				t.Fatalf("неожиданный id: %d", id)
			}

			if info.Name != input.Name || info.Email != input.Email {
				t.Fatalf("неожиданные данные пользователя: %+v", info)
			}

			if info.Sex != "female" {
				t.Fatalf("ожидался нормализованный пол, получено %q", info.Sex)
			}

			if info.TargetCalories != 1329 || info.TargetProteins != 110 || info.TargetFats != 41 || info.TargetCarbs != 131 {
				t.Fatalf("неожиданные целевые значения: %+v", info)
			}
		}).
		Return(nil)

	if err := service.Update(ctx, 7, input); err != nil {
		t.Fatalf("Update вернул ошибку: %v", err)
	}

	userStorage.AssertExpectations(t)
}

func TestServiceUpdate_ReturnsValidationError(t *testing.T) {
	service := newServiceForTest(&mocks.UserStorage{}, nil)
	input := *validUserInput()
	input.Email = "bad-email"

	err := service.Update(context.Background(), 1, input)
	if err == nil {
		t.Fatalf("ожидалась ошибка валидации")
	}

	if !strings.Contains(err.Error(), "некорректный email") {
		t.Fatalf("неожиданная ошибка: %v", err)
	}
}

func TestServiceUpdate_WrapsCalculationError(t *testing.T) {
	service := newPermissiveServiceForTest(&mocks.UserStorage{}, nil)

	err := service.Update(context.Background(), 1, models.UserInput{
		Name:           "A",
		Email:          "a@example.com",
		Sex:            "male",
		Age:            0,
		HeightCm:       0,
		WeightKg:       0,
		TargetWeightKg: 0,
	})
	if err == nil {
		t.Fatalf("ожидалась ошибка расчета")
	}

	if !strings.Contains(err.Error(), "не удалось пересчитать цели") {
		t.Fatalf("неожиданная ошибка: %v", err)
	}
}
