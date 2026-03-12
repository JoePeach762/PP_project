package user

import (
	"context"
	"time"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
)

type userStorage interface {
	AddUsers(ctx context.Context, infos []*models.UserInfo) error
	GetUsersByIds(ctx context.Context, ids []uint64) ([]*models.UserInfo, error)
	UpdateUser(ctx context.Context, id uint64, info models.UserInfo) error
	DeleteUsers(ctx context.Context, ids []uint64) error
}

type statsStorage interface {
	AddMealToUser(ctx context.Context, mealInfo *models.MealInfo) error
	GetStatsByUserIDs(ctx context.Context, ids []uint64, date time.Time) ([]*models.UserStats, error)
}

type producer interface {
	PublishUsersDeleted(ctx context.Context, ids []uint64) error
}

type Service struct {
	userStorage   userStorage
	statsStorage  statsStorage
	producer      producer
	minNameLength uint32
	maxNameLength uint32
	minWeight     uint32
	maxWeight     uint32
}

func NewUserService(
	userStorage userStorage,
	statsStorage statsStorage,
	producer producer,
	minNameLength uint32,
	maxNameLength uint32,
	minWeight uint32,
	maxWeight uint32,
) *Service {
	return &Service{
		userStorage:   userStorage,
		statsStorage:  statsStorage,
		producer:      producer,
		minNameLength: minNameLength,
		maxNameLength: maxNameLength,
		minWeight:     minWeight,
		maxWeight:     maxWeight,
	}
}
