package outbox

import (
	"context"

	"openhealth/internal/event"
)

//go:generate mockgen -source=interface.go -destination=test/mocks.go -package=test

type Repository interface {
	AddOutboxEvent(ctx context.Context, outboxEvent event.OutboxEvent) error
	UpdateOutboxEventStatus(ctx context.Context, id string, status event.OutboxEventStatus) error
	GetPendingOutboxEvents(ctx context.Context, limit int) ([]event.OutboxEvent, error)
}
