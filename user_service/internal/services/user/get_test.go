package user

import (
	"context"
	"errors"
	"testing"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
	"github.com/JoePeach762/PP_project/user_service/internal/services/user/mocks"
	"github.com/stretchr/testify/mock"
)

func TestServiceGetByIds_EmptyInput(t *testing.T) {
	service := newServiceForTest(&mocks.UserStorage{}, &mocks.StatsStorage{})

	users, err := service.GetByIds(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetByIds вернул ошибку: %v", err)
	}

	if len(users) != 0 {
		t.Fatalf("ожидался пустой результат, получено пользователей: %d", len(users))
	}
}

func TestServiceGetByIds_ReturnsUserStorageError(t *testing.T) {
	ctx := context.Background()
	userStorage := &mocks.UserStorage{}
	service := newServiceForTest(userStorage, &mocks.StatsStorage{})
	expectedErr := errors.New("ошибка получения пользователей")

	userStorage.EXPECT().GetUsersByIds(ctx, []uint64{1, 2}).Return(nil, expectedErr)

	_, err := service.GetByIds(ctx, []uint64{1, 2})
	if !errors.Is(err, expectedErr) {
		t.Fatalf("ожидалась ошибка %v, получено %v", expectedErr, err)
	}
}

func TestServiceGetByIds_ReturnsStatsStorageError(t *testing.T) {
	ctx := context.Background()
	userStorage := &mocks.UserStorage{}
	statsStorage := &mocks.StatsStorage{}
	service := newServiceForTest(userStorage, statsStorage)
	expectedErr := errors.New("ошибка получения статистики")

	userStorage.EXPECT().GetUsersByIds(ctx, []uint64{1}).Return([]*models.UserInfo{{ID: 1, Name: "Alice"}}, nil)
	statsStorage.EXPECT().GetStatsByUserIDs(ctx, []uint64{1}, mock.Anything).Return(nil, expectedErr)

	_, err := service.GetByIds(ctx, []uint64{1})
	if !errors.Is(err, expectedErr) {
		t.Fatalf("ожидалась ошибка %v, получено %v", expectedErr, err)
	}
}

func TestServiceGetByIds_MergesUsersAndStats(t *testing.T) {
	ctx := context.Background()
	userStorage := &mocks.UserStorage{}
	statsStorage := &mocks.StatsStorage{}
	service := newServiceForTest(userStorage, statsStorage)
	ids := []uint64{1, 2}

	userStorage.EXPECT().GetUsersByIds(ctx, ids).Return([]*models.UserInfo{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Sex: "female", Age: 28, HeightCm: 165, WeightKg: 60, TargetWeightKg: 55, TargetCalories: 1329, TargetProteins: 110, TargetFats: 41, TargetCarbs: 131},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Sex: "male", Age: 30, HeightCm: 180, WeightKg: 80, TargetWeightKg: 75, TargetCalories: 1948, TargetProteins: 150, TargetFats: 60, TargetCarbs: 203},
	}, nil)
	statsStorage.EXPECT().GetStatsByUserIDs(ctx, ids, mock.Anything).Return([]*models.UserStats{
		{UserID: 1, CurrentCalories: 800, CurrentProteins: 70, CurrentFats: 20, CurrentCarbs: 90},
	}, nil)

	users, err := service.GetByIds(ctx, ids)
	if err != nil {
		t.Fatalf("GetByIds вернул ошибку: %v", err)
	}

	if len(users) != 2 {
		t.Fatalf("неожиданное количество пользователей: %d", len(users))
	}

	if users[0].CurrentCalories != 800 || users[0].CurrentProteins != 70 {
		t.Fatalf("ожидалось объединение статистики для первого пользователя, получено %+v", users[0])
	}

	if users[1].CurrentCalories != 0 || users[1].CurrentProteins != 0 {
		t.Fatalf("для второго пользователя ожидалась нулевая статистика, получено %+v", users[1])
	}
}
