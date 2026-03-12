package mealstorage

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
		return nil, errors.Wrap(err, "generate meal outbox SELECT query")
	}

	rows, err := s.db.Query(ctx, queryText, args...)
	if err != nil {
		return nil, errors.Wrap(err, "execute meal outbox SELECT query")
	}
	defer rows.Close()

	events := make([]OutboxEvent, 0, limit)
	for rows.Next() {
		var event OutboxEvent
		if err := rows.Scan(&event.ID, &event.Payload); err != nil {
			return nil, errors.Wrap(err, "scan meal outbox row")
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
		return errors.Wrap(err, "generate meal outbox UPDATE query")
	}

	if _, err := s.db.Exec(ctx, queryText, args...); err != nil {
		return errors.Wrap(err, "execute meal outbox UPDATE query")
	}

	return nil
}
