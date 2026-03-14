package mealoutbox

import (
	"context"
	"log/slog"
	"time"

	mealstorage "github.com/JoePeach762/PP_project/meal_service/internal/storage/pgstorage/mealstorage"
)

const (
	defaultBatchSize    = 100
	defaultPollInterval = time.Second
)

type storage interface {
	FetchPendingOutboxEvents(ctx context.Context, limit int) ([]mealstorage.OutboxEvent, error)
	MarkOutboxEventPublished(ctx context.Context, id uint64) error
}

type producer interface {
	PublishMessage(ctx context.Context, payload []byte) error
}

type Publisher struct {
	storage      storage
	producer     producer
	batchSize    int
	pollInterval time.Duration
}

func NewPublisher(storage storage, producer producer) *Publisher {
	return &Publisher{
		storage:      storage,
		producer:     producer,
		batchSize:    defaultBatchSize,
		pollInterval: defaultPollInterval,
	}
}

func (p *Publisher) Run(ctx context.Context) {
	ticker := time.NewTicker(p.pollInterval)
	defer ticker.Stop()

	p.publishPending(ctx)

	for {
		select {
		case <-ctx.Done():
			slog.Info("Публикатор meal outbox остановлен")
			return
		case <-ticker.C:
			p.publishPending(ctx)
		}
	}
}

func (p *Publisher) publishPending(ctx context.Context) {
	for {
		if ctx.Err() != nil {
			return
		}

		events, err := p.storage.FetchPendingOutboxEvents(ctx, p.batchSize)
		if err != nil {
			slog.Error("Не удалось получить события из meal outbox", "error", err)
			return
		}
		if len(events) == 0 {
			return
		}

		for _, event := range events {
			if err := p.producer.PublishMessage(ctx, event.Payload); err != nil {
				slog.Error("Не удалось опубликовать событие из meal outbox", "error", err, "event_id", event.ID)
				return
			}

			if err := p.storage.MarkOutboxEventPublished(ctx, event.ID); err != nil {
				slog.Error("Не удалось отметить событие из meal outbox как опубликованное", "error", err, "event_id", event.ID)
				return
			}
		}

		if len(events) < p.batchSize {
			return
		}
	}
}
