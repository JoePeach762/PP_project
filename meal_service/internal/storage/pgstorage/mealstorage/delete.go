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
	for _, id := range ids {
		if err := s.deleteSingle(ctx, id); err != nil {
			return err
		}
	}

	return nil
}

func (s *PGstorage) deleteSingle(ctx context.Context, id uint64) error {
	query := squirrel.Delete(mealTableName).
		Where(squirrel.Eq{mealUserIDcolumnName: id}).
		PlaceholderFormat(squirrel.Dollar)

	queryText, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "generate delete !meals! query")
	}

	_, err = s.db.Exec(ctx, queryText, args...)
	if err != nil {
		return errors.Wrap(err, "execute delete !meals! query")
	}

	return nil
}
