package outbox

import (
	"context"

	"openhealth/internal/event"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// AddOutboxEvent - adds a new outbox event
func (s *Service) AddOutboxEvent(ctx context.Context, outboxEvent event.OutboxEvent) error {
	return s.repository.AddOutboxEvent(ctx, outboxEvent)
}

// UpdateOutboxEventStatus - updates the status of an outbox event
func (s *Service) UpdateOutboxEventStatus(ctx context.Context, id string, status event.OutboxEventStatus) error {
	return s.repository.UpdateOutboxEventStatus(ctx, id, status)
}

// GetPendingOutboxEvents - gets pending outbox events up to the provided limit
func (s *Service) GetPendingOutboxEvents(ctx context.Context, limit int) ([]event.OutboxEvent, error) {
	return s.repository.GetPendingOutboxEvents(ctx, limit)
}
