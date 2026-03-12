package user

import (
	"context"
	"time"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
)

func (s *Service) GetByIds(ctx context.Context, ids []uint64) ([]*models.UserWithStats, error) {
	if len(ids) == 0 {
		return []*models.UserWithStats{}, nil
	}

	users, err := s.userStorage.GetUsersByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	stats, err := s.statsStorage.GetStatsByUserIDs(ctx, ids, time.Now())
	if err != nil {
		return nil, err
	}

	statsByUserID := make(map[uint64]*models.UserStats, len(stats))
	for _, stat := range stats {
		statsByUserID[stat.UserID] = stat
	}

	usersWithStats := make([]*models.UserWithStats, 0, len(users))
	for _, user := range users {
		usersWithStats = append(usersWithStats, models.NewUserWithStats(user, statsByUserID[user.ID]))
	}

	return usersWithStats, nil
}
