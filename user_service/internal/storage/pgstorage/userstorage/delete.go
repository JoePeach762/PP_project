package userstorage

import (
	"context"
	"encoding/json"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
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

func (s *PGstorage) DeleteUsersAndEnqueueEvent(ctx context.Context, ids []uint64) (err error) {
	if len(ids) == 0 {
		return nil
	}

	payload, err := json.Marshal(ids)
	if err != nil {
		return errors.Wrap(err, "marshal deleted users event")
	}

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return errors.Wrap(err, "begin delete users transaction")
	}
	defer func() {
		if err == nil {
			return
		}
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil && rollbackErr != pgx.ErrTxClosed {
			err = errors.Wrap(rollbackErr, "rollback delete users transaction")
		}
	}()

	deleteQuery := squirrel.Delete(userTableName).
		Where(squirrel.Eq{userIDColumnName: ids}).
		PlaceholderFormat(squirrel.Dollar)

	deleteQueryText, deleteArgs, err := deleteQuery.ToSql()
	if err != nil {
		return errors.Wrap(err, "generate delete users query")
	}
	if _, err = tx.Exec(ctx, deleteQueryText, deleteArgs...); err != nil {
		return errors.Wrap(err, "execute delete users query")
	}

	outboxQuery := squirrel.Insert(outboxTableName).
		Columns(outboxPayloadColumnName).
		Values(payload).
		PlaceholderFormat(squirrel.Dollar)

	outboxQueryText, outboxArgs, err := outboxQuery.ToSql()
	if err != nil {
		return errors.Wrap(err, "generate user outbox query")
	}
	if _, err = tx.Exec(ctx, outboxQueryText, outboxArgs...); err != nil {
		return errors.Wrap(err, "execute user outbox query")
	}

	if err = tx.Commit(ctx); err != nil {
		return errors.Wrap(err, "commit delete users transaction")
	}

	return nil
}
