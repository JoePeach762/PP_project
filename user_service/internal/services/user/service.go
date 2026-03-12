package user

import (
	"context"
	"strings"
	"time"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
)

type userStorage interface {
	AddUsers(ctx context.Context, infos []*models.UserInfo) error
	GetUsersByIds(ctx context.Context, ids []uint64) ([]*models.UserInfo, error)
	UpdateUser(ctx context.Context, id uint64, info models.UserInfo) error
	DeleteUsers(ctx context.Context, ids []uint64) error
	DeleteUsersAndEnqueueEvent(ctx context.Context, ids []uint64) error
}

type statsStorage interface {
	AddMealToUser(ctx context.Context, mealInfo *models.MealInfo) error
	GetStatsByUserIDs(ctx context.Context, ids []uint64, date time.Time) ([]*models.UserStats, error)
}

type Service struct {
	userStorage   userStorage
	statsStorage  statsStorage
	minNameLength uint32
	maxNameLength uint32
	minAge        uint32
	maxAge        uint32
	minHeightCm   uint32
	maxHeightCm   uint32
	minWeight     uint32
	maxWeight     uint32
	allowedSexes  map[string]struct{}
}

func NewUserService(
	userStorage userStorage,
	statsStorage statsStorage,
	minNameLength uint32,
	maxNameLength uint32,
	minAge uint32,
	maxAge uint32,
	minHeightCm uint32,
	maxHeightCm uint32,
	minWeight uint32,
	maxWeight uint32,
	allowedSexes []string,
) *Service {
	allowedSexesMap := make(map[string]struct{}, len(allowedSexes))
	for _, sex := range allowedSexes {
		allowedSexesMap[strings.ToLower(strings.TrimSpace(sex))] = struct{}{}
	}

	return &Service{
		userStorage:   userStorage,
		statsStorage:  statsStorage,
		minNameLength: minNameLength,
		maxNameLength: maxNameLength,
		minAge:        minAge,
		maxAge:        maxAge,
		minHeightCm:   minHeightCm,
		maxHeightCm:   maxHeightCm,
		minWeight:     minWeight,
		maxWeight:     maxWeight,
		allowedSexes:  allowedSexesMap,
	}
}
