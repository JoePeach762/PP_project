// internal/storage/pgstorage/delete.go
package pgstorage

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (s *PGstorage) DeleteUsers(ctx context.Context, ids []uint64) error {
	if len(ids) == 0 {
		return nil
	}

	query := squirrel.Delete(userTableName).
		Where(squirrel.Eq{userIDColumnName: ids}).
		PlaceholderFormat(squirrel.Dollar)

	queryText, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "generate delete !users! query")
	}

	_, err = s.db.Exec(ctx, queryText, args...)
	if err != nil {
		return errors.Wrap(err, "execute delete !users! query")
	}

	return nil
}
