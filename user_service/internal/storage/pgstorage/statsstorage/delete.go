package statsstorage

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (s *PGstorage) DeleteByUserIDs(ctx context.Context, ids []uint64) error {
	if len(ids) == 0 {
		return nil
	}

	query := squirrel.Delete(statsTableName).
		Where(squirrel.Eq{statsUserIDColumnName: ids}).
		PlaceholderFormat(squirrel.Dollar)

	queryText, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "не удалось сформировать запрос на удаление статистики")
	}

	_, err = s.db.Exec(ctx, queryText, args...)
	if err != nil {
		return errors.Wrap(err, "не удалось выполнить запрос на удаление статистики")
	}

	return nil
}
