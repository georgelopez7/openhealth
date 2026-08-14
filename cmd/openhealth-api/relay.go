package main

import (
	"context"
	"log/slog"
	"time"

	"openhealth/internal/event"
	"openhealth/internal/service/outbox"
)

const relayBatchSize = 100

type Relay struct {
	outboxService *outbox.Service
	mutex         *event.EventMutex
	done          chan struct{}
}

// NewRelay - creates a new outbox relay
func NewRelay(outboxService *outbox.Service) *Relay {
	return &Relay{
		outboxService: outboxService,
		mutex:         event.NewEventMutex(),
		done:          make(chan struct{}),
	}
}

// Start - begins polling for pending outbox events on a 2 second timer
func (r *Relay) Start(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			r.processPendingEvents(ctx)
		case <-r.done:
			return
		case <-ctx.Done():
			return
		}
	}
}

// Stop - stops the relay
func (r *Relay) Stop() {
	close(r.done)
}

func (r *Relay) processPendingEvents(ctx context.Context) {
	outboxEvents, err := r.outboxService.GetPendingOutboxEvents(ctx, relayBatchSize)
	if err != nil {
		slog.Error("failed to get pending outbox events", "err", err)
		return
	}

	for _, outboxEvent := range outboxEvents {
		if !r.mutex.Add(outboxEvent.EventID) {
			continue
		}

		go r.handleEvent(ctx, outboxEvent)
	}
}

func (r *Relay) handleEvent(ctx context.Context, outboxEvent event.OutboxEvent) {
	defer r.mutex.Remove(outboxEvent.EventID)

	switch outboxEvent.EventType {
	case event.EventTypeAccountCreated:
	// TODO: handle account created
	case event.EventTypeAccountUpdated:
	// TODO: handle account updated
	case event.EventTypeAccountDeleted:
	// TODO: handle account deleted
	case event.EventTypeDoctorCreated:
	// TODO: handle doctor created
	case event.EventTypeDoctorUpdated:
	// TODO: handle doctor updated
	case event.EventTypeDoctorDeleted:
	// TODO: handle doctor deleted
	case event.EventTypeHospitalCreated:
	// TODO: handle hospital created
	case event.EventTypeHospitalUpdated:
	// TODO: handle hospital updated
	case event.EventTypeHospitalDeleted:
	// TODO: handle hospital deleted
	case event.EventTypeNurseCreated:
	// TODO: handle nurse created
	case event.EventTypeNurseUpdated:
	// TODO: handle nurse updated
	case event.EventTypeNurseDeleted:
	// TODO: handle nurse deleted
	default:
		slog.Error("unknown event type", "event_type", outboxEvent.EventType, "event", outboxEvent.EventJSON)
		return
	}

	err := r.outboxService.UpdateOutboxEventStatus(ctx, outboxEvent.ID, event.OutboxEventStatusProcessed)
	if err != nil {
		slog.Error("failed to update event status to processed", "id", outboxEvent.ID, "event", outboxEvent.EventJSON)
	}
}
