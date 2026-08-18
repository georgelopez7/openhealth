package test

import (
	"errors"
	"openhealth/internal/domain"
	"openhealth/internal/event"
	"testing"

	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

func TestService_AddOutboxEvent(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully add outbox event", func(t *testing.T) {
		account := domain.NewAccount("John", "Doe", 30, "john.doe@example.com", "")
		accountEvent := event.NewAccountCreatedEvent(*account)

		outboxEvent, err := event.NewOutboxEvent(accountEvent)
		require.NoError(t, err)

		deps.repository.EXPECT().AddOutboxEvent(gomock.Any(), *outboxEvent).Return(nil)

		err = service.AddOutboxEvent(ctx, *outboxEvent)
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails", func(t *testing.T) {
		account := domain.NewAccount("John", "Doe", 30, "john.doe@example.com", "")
		accountEvent := event.NewAccountCreatedEvent(*account)

		outboxEvent, err := event.NewOutboxEvent(accountEvent)
		require.NoError(t, err)

		deps.repository.EXPECT().AddOutboxEvent(gomock.Any(), *outboxEvent).Return(errors.New("add error"))

		err = service.AddOutboxEvent(ctx, *outboxEvent)
		require.Error(t, err)
	})
}

func TestService_UpdateOutboxEventStatus(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully update outbox event status", func(t *testing.T) {
		deps.repository.EXPECT().UpdateOutboxEventStatus(gomock.Any(), "outbox-event-id", event.OutboxEventStatusProcessed).Return(nil)

		err := service.UpdateOutboxEventStatus(ctx, "outbox-event-id", event.OutboxEventStatusProcessed)
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails", func(t *testing.T) {
		deps.repository.EXPECT().UpdateOutboxEventStatus(gomock.Any(), "outbox-event-id", event.OutboxEventStatusFailed).Return(errors.New("update error"))

		err := service.UpdateOutboxEventStatus(ctx, "outbox-event-id", event.OutboxEventStatusFailed)
		require.Error(t, err)
	})
}

func TestService_GetPendingOutboxEvents(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully return pending outbox events", func(t *testing.T) {
		account := domain.NewAccount("John", "Doe", 30, "john.doe@example.com", "")
		accountEvent := event.NewAccountCreatedEvent(*account)

		outboxEvent, err := event.NewOutboxEvent(accountEvent)
		require.NoError(t, err)

		expected := []event.OutboxEvent{*outboxEvent}

		deps.repository.EXPECT().GetPendingOutboxEvents(gomock.Any(), 10).Return(expected, nil)

		pendingEvents, err := service.GetPendingOutboxEvents(ctx, 10)
		require.NoError(t, err)
		require.Len(t, pendingEvents, 1)
		require.Equal(t, outboxEvent.ID, pendingEvents[0].ID)
	})

	t.Run("should handle error when repository fails", func(t *testing.T) {
		deps.repository.EXPECT().GetPendingOutboxEvents(gomock.Any(), 10).Return(nil, errors.New("get error"))

		pendingEvents, err := service.GetPendingOutboxEvents(ctx, 10)
		require.Error(t, err)
		require.Nil(t, pendingEvents)
	})
}
