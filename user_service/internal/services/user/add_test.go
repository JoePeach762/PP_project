package user

import (
	"context"
	"strings"
	"testing"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
	"github.com/JoePeach762/PP_project/user_service/internal/services/user/mocks"
	"github.com/stretchr/testify/mock"
)

func TestServiceAdd_SavesCalculatedUsers(t *testing.T) {
	ctx := context.Background()
	userStorage := &mocks.UserStorage{}
	service := newServiceForTest(userStorage, nil)
	input := validUserInput()

	userStorage.EXPECT().
		AddUsers(ctx, mock.Anything).
		Run(func(_ context.Context, infos []*models.UserInfo) {
			if len(infos) != 1 {
				t.Fatalf("unexpected infos length: %d", len(infos))
			}

			info := infos[0]
			if info.Name != input.Name || info.Email != input.Email {
				t.Fatalf("unexpected user info: %+v", info)
			}

			if info.Sex != "female" {
				t.Fatalf("expected normalized sex, got %q", info.Sex)
			}

			if info.TargetCalories != 1329 || info.TargetProteins != 110 || info.TargetFats != 41 || info.TargetCarbs != 131 {
				t.Fatalf("unexpected targets: %+v", info)
			}
		}).
		Return(nil)

	if err := service.Add(ctx, []*models.UserInput{input}); err != nil {
		t.Fatalf("Add returned error: %v", err)
	}

	userStorage.AssertExpectations(t)
}

func TestServiceAdd_ReturnsValidationError(t *testing.T) {
	service := newServiceForTest(&mocks.UserStorage{}, nil)
	input := validUserInput()
	input.Email = "bad-email"

	err := service.Add(context.Background(), []*models.UserInput{input})
	if err == nil {
		t.Fatalf("expected validation error")
	}

	if !strings.Contains(err.Error(), "некорректный email") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestServiceAdd_WrapsCalculationError(t *testing.T) {
	service := newPermissiveServiceForTest(&mocks.UserStorage{}, nil)

	err := service.Add(context.Background(), []*models.UserInput{{
		Name:           "A",
		Email:          "a@example.com",
		Sex:            "male",
		Age:            0,
		HeightCm:       0,
		WeightKg:       0,
		TargetWeightKg: 0,
	}})
	if err == nil {
		t.Fatalf("expected calculation error")
	}

	if !strings.Contains(err.Error(), "не удалось рассчитать цели") {
		t.Fatalf("unexpected error: %v", err)
	}
}
