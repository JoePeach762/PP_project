package statsstorage

import (
	"context"
	"time"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (s *PGstorage) GetStatsByUserIDs(ctx context.Context, ids []uint64, date time.Time) ([]*models.UserStats, error) {
	if len(ids) == 0 {
		return []*models.UserStats{}, nil
	}

	query := squirrel.Select(
		statsUserIDColumnName,
		statsCaloriesColumnName,
		statsProteinsColumnName,
		statsFatsColumnName,
		statsCarbsColumnName,
		statsDateColumnName,
	).
		From(statsTableName).
		Where(squirrel.Eq{statsUserIDColumnName: ids}).
		Where(squirrel.Eq{statsDateColumnName: normalizeDate(date)}).
		PlaceholderFormat(squirrel.Dollar)

	queryText, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "generate !stats! query error")
	}

	rows, err := s.db.Query(ctx, queryText, args...)
	if err != nil {
		return nil, errors.Wrap(err, "quering !stats! error")
	}
	defer rows.Close()

	var stats []*models.UserStats
	for rows.Next() {
		var stat models.UserStats
		if err := rows.Scan(
			&stat.UserID,
			&stat.CurrentCalories,
			&stat.CurrentProteins,
			&stat.CurrentFats,
			&stat.CurrentCarbs,
			&stat.Date,
		); err != nil {
			return nil, errors.Wrap(err, "failed to scan !stats! row")
		}
		stats = append(stats, &stat)
	}

	return stats, nil
}
