package mealstorage

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (s *PGstorage) DeleteByUserIds(ctx context.Context, ids []uint64) error {
	if len(ids) == 0 {
		return nil
	}
	query := squirrel.Delete(mealTableName).
		Where(squirrel.Eq{mealUserIDcolumnName: ids}).
		PlaceholderFormat(squirrel.Dollar)

	queryText, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "не удалось сформировать запрос на удаление приемов пищи")
	}

	_, err = s.db.Exec(ctx, queryText, args...)
	if err != nil {
		return errors.Wrap(err, "не удалось выполнить запрос на удаление приемов пищи")
	}

	return nil
}
