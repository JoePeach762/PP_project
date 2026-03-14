package userstorage

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (s *PGstorage) FetchPendingOutboxEvents(ctx context.Context, limit int) ([]OutboxEvent, error) {
	query := squirrel.Select(outboxIDColumnName, outboxPayloadColumnName).
		From(outboxTableName).
		Where(squirrel.Expr(outboxPublishedAtColumnName + " IS NULL")).
		OrderBy(outboxIDColumnName).
		Limit(uint64(limit)).
		PlaceholderFormat(squirrel.Dollar)

	queryText, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "не удалось сформировать SELECT-запрос к user outbox")
	}

	rows, err := s.db.Query(ctx, queryText, args...)
	if err != nil {
		return nil, errors.Wrap(err, "не удалось выполнить SELECT-запрос к user outbox")
	}
	defer rows.Close()

	events := make([]OutboxEvent, 0, limit)
	for rows.Next() {
		var event OutboxEvent
		if err := rows.Scan(&event.ID, &event.Payload); err != nil {
			return nil, errors.Wrap(err, "не удалось прочитать строку user outbox")
		}
		events = append(events, event)
	}

	return events, nil
}

func (s *PGstorage) MarkOutboxEventPublished(ctx context.Context, id uint64) error {
	query := squirrel.Update(outboxTableName).
		Set(outboxPublishedAtColumnName, squirrel.Expr("NOW()")).
		Where(squirrel.Eq{outboxIDColumnName: id}).
		Where(squirrel.Expr(outboxPublishedAtColumnName + " IS NULL")).
		PlaceholderFormat(squirrel.Dollar)

	queryText, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "не удалось сформировать UPDATE-запрос к user outbox")
	}

	if _, err := s.db.Exec(ctx, queryText, args...); err != nil {
		return errors.Wrap(err, "не удалось выполнить UPDATE-запрос к user outbox")
	}

	return nil
}
