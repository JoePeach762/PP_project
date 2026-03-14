package user

import (
	"context"
	"errors"
	"testing"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
	"github.com/JoePeach762/PP_project/user_service/internal/services/user/mocks"
)

func TestServiceAddMealToUser_DelegatesToStatsStorage(t *testing.T) {
	ctx := context.Background()
	statsStorage := &mocks.StatsStorage{}
	service := newServiceForTest(nil, statsStorage)
	meal := &models.MealInfo{UserId: 42, Name: "Apple", WeightGrams: 150}
	expectedErr := errors.New("ошибка хранилища")

	statsStorage.EXPECT().AddMealToUser(ctx, meal).Return(expectedErr)

	err := service.AddMealToUser(ctx, meal)
	if !errors.Is(err, expectedErr) {
		t.Fatalf("ожидалась ошибка %v, получено %v", expectedErr, err)
	}

	statsStorage.AssertExpectations(t)
}
