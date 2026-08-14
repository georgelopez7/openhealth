package test

import (
	"openhealth/internal/domain"
	"openhealth/internal/event"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepository_AddOutboxEvent(t *testing.T) {
	ctx := t.Context()

	repo.ResetOutboxEvents(ctx)

	account := domain.NewAccount("John", "Doe", 30, "john.doe@example.com")
	accountEvent := event.NewAccountCreatedEvent(*account)

	expected, err := event.NewOutboxEvent(accountEvent)
	require.NoError(t, err)

	err = repo.AddOutboxEvent(ctx, *expected)
	require.NoError(t, err)

	pendingEvents, err := repo.GetPendingOutboxEvents(ctx, 10)
	require.NoError(t, err)
	require.Len(t, pendingEvents, 1)
	require.Equal(t, expected.ID, pendingEvents[0].ID)
	require.Equal(t, expected.EventID, pendingEvents[0].EventID)
	require.Equal(t, expected.EventType, pendingEvents[0].EventType)
	require.Equal(t, event.OutboxEventStatusPending, pendingEvents[0].Status)
	require.JSONEq(t, string(expected.EventJSON), string(pendingEvents[0].EventJSON))
}

func TestRepository_UpdateOutboxEventStatus(t *testing.T) {
	ctx := t.Context()

	repo.ResetOutboxEvents(ctx)

	account := domain.NewAccount("John", "Doe", 30, "john.doe@example.com")
	accountEvent := event.NewAccountCreatedEvent(*account)

	outboxEvent, err := event.NewOutboxEvent(accountEvent)
	require.NoError(t, err)

	err = repo.AddOutboxEvent(ctx, *outboxEvent)
	require.NoError(t, err)

	err = repo.UpdateOutboxEventStatus(ctx, outboxEvent.ID, event.OutboxEventStatusProcessed)
	require.NoError(t, err)

	pendingEvents, err := repo.GetPendingOutboxEvents(ctx, 10)
	require.NoError(t, err)
	require.Empty(t, pendingEvents)
}

func TestRepository_GetPendingOutboxEvents(t *testing.T) {
	ctx := t.Context()

	repo.ResetOutboxEvents(ctx)

	firstAccount := domain.NewAccount("John", "Doe", 30, "john.doe@example.com")
	firstEvent := event.NewAccountCreatedEvent(*firstAccount)
	first, err := event.NewOutboxEvent(firstEvent)
	require.NoError(t, err)

	secondAccount := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com")
	secondEvent := event.NewAccountCreatedEvent(*secondAccount)
	second, err := event.NewOutboxEvent(secondEvent)
	require.NoError(t, err)

	err = repo.AddOutboxEvent(ctx, *first)
	require.NoError(t, err)

	err = repo.AddOutboxEvent(ctx, *second)
	require.NoError(t, err)

	t.Run("should return pending events up to the limit ordered by created_at", func(t *testing.T) {
		pendingEvents, err := repo.GetPendingOutboxEvents(ctx, 1)
		require.NoError(t, err)
		require.Len(t, pendingEvents, 1)
		require.Equal(t, first.ID, pendingEvents[0].ID)
	})

	t.Run("should return all pending events when limit is greater than count", func(t *testing.T) {
		pendingEvents, err := repo.GetPendingOutboxEvents(ctx, 10)
		require.NoError(t, err)
		require.Len(t, pendingEvents, 2)
		require.Equal(t, first.ID, pendingEvents[0].ID)
		require.Equal(t, second.ID, pendingEvents[1].ID)
	})

	t.Run("should not return processed events", func(t *testing.T) {
		err := repo.UpdateOutboxEventStatus(ctx, first.ID, event.OutboxEventStatusProcessed)
		require.NoError(t, err)

		pendingEvents, err := repo.GetPendingOutboxEvents(ctx, 10)
		require.NoError(t, err)
		require.Len(t, pendingEvents, 1)
		require.Equal(t, second.ID, pendingEvents[0].ID)
	})

	t.Run("should return empty slice when no pending events exist", func(t *testing.T) {
		err := repo.UpdateOutboxEventStatus(ctx, second.ID, event.OutboxEventStatusProcessed)
		require.NoError(t, err)

		pendingEvents, err := repo.GetPendingOutboxEvents(ctx, 10)
		require.NoError(t, err)
		require.Empty(t, pendingEvents)
	})
}
