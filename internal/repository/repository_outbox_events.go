package repository

import (
	"context"

	"openhealth/internal/event"
	"openhealth/internal/pkg/postgres"
)

// AddOutboxEvent - adds a new outbox event to the database
func (r *Repository) AddOutboxEvent(ctx context.Context, outboxEvent event.OutboxEvent) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		INSERT INTO outbox_events (id, status, event_id, event_type, event_json, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, outboxEvent.ID, outboxEvent.Status, outboxEvent.EventID, outboxEvent.EventType, string(outboxEvent.EventJSON), outboxEvent.CreatedAt)

	return err
}

// UpdateOutboxEventStatus - updates the status of an outbox event
func (r *Repository) UpdateOutboxEventStatus(ctx context.Context, id string, status event.OutboxEventStatus) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		UPDATE outbox_events
		SET status = $1
		WHERE id = $2
	`, status, id)

	return err
}

// GetPendingOutboxEvents - gets pending outbox events up to the provided limit, ordered by creation time
func (r *Repository) GetPendingOutboxEvents(ctx context.Context, limit int) ([]event.OutboxEvent, error) {
	var outboxEvents []event.OutboxEvent

	err := r.db.SelectContext(ctx, &outboxEvents, `
		SELECT id, status, event_id, event_type, event_json, created_at
		FROM outbox_events
		WHERE status = $1
		ORDER BY created_at ASC
		LIMIT $2
	`, event.OutboxEventStatusPending, limit)

	return outboxEvents, err
}

// ResetOutboxEvents - resets all outbox events
func (r *Repository) ResetOutboxEvents(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `
		TRUNCATE TABLE outbox_events
	`)

	return err
}
