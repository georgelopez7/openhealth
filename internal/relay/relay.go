package relay

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"openhealth/internal/event"
	"openhealth/internal/service/authorization"
	"openhealth/internal/service/outbox"
)

type Relay struct {
	outboxService        *outbox.Service
	authorizationService *authorization.Service
	batchSize            int
	tick                 time.Duration
	mutex                *event.EventMutex
	done                 chan struct{}
}

// NewRelay - creates a new outbox relay
func NewRelay(outboxService *outbox.Service, authorizationService *authorization.Service, batchSize int, tick time.Duration) *Relay {
	return &Relay{
		outboxService:        outboxService,
		authorizationService: authorizationService,
		batchSize:            batchSize,
		tick:                 tick,
		mutex:                event.NewEventMutex(),
		done:                 make(chan struct{}),
	}
}

// Start - begins polling for pending outbox events on a timer
func (r *Relay) Start(ctx context.Context) {
	ticker := time.NewTicker(r.tick)
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

// processPendingEvents - processes pending outbox events
func (r *Relay) processPendingEvents(ctx context.Context) {
	outboxEvents, err := r.outboxService.GetPendingOutboxEvents(ctx, r.batchSize)
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

// handleEvent - handles an outbox event
func (r *Relay) handleEvent(ctx context.Context, outboxEvent event.OutboxEvent) {
	defer r.mutex.Remove(outboxEvent.EventID)

	var _event event.Event
	if err := json.Unmarshal(outboxEvent.EventJSON, &_event); err != nil {
		slog.Error("failed to unmarshal outbox event", "outbox_event_id", outboxEvent.ID, "event", outboxEvent.EventJSON, "err", err)
		return
	}

	var handleErr error
	switch _event.Type {
	case event.EventTypeAccountCreated:
		// skip - account.created
	case event.EventTypeAccountUpdated:
		// skip - account.updated
	case event.EventTypeAccountDeleted:
		// skip - account.deleted
	case event.EventTypeDoctorCreated:
		handleErr = r.handleDoctorCreated(ctx, _event)
	case event.EventTypeDoctorUpdated:
		// skip - doctor.updated
	case event.EventTypeDoctorDeleted:
		handleErr = r.handleDoctorDeleted(ctx, _event)
	case event.EventTypeHospitalCreated:
		// skip - hospital.created
	case event.EventTypeHospitalUpdated:
		// skip - hospital.updated
	case event.EventTypeHospitalDeleted:
		// skip - hospital.deleted
	case event.EventTypeNurseCreated:
		handleErr = r.handleNurseCreated(ctx, _event)
	case event.EventTypeNurseUpdated:
		// skip - nurse.updated
	case event.EventTypeNurseDeleted:
		handleErr = r.handleNurseDeleted(ctx, _event)
	case event.EventTypeDoctorToAccountAssignmentCreated:
		handleErr = r.handleDoctorToAccountAssignmentCreated(ctx, _event)
	case event.EventTypeDoctorToAccountAssignmentRemoved:
		handleErr = r.handleDoctorToAccountAssignmentRemoved(ctx, _event)
	case event.EventTypeHospitalToAccountAssignmentCreated:
		handleErr = r.handleHospitalToAccountAssignmentCreated(ctx, _event)
	case event.EventTypeHospitalToAccountAssignmentRemoved:
		handleErr = r.handleHospitalToAccountAssignmentRemoved(ctx, _event)
	case event.EventTypeNurseToHospitalAssignmentCreated:
		handleErr = r.handleNurseToHospitalAssignmentCreated(ctx, _event)
	case event.EventTypeNurseToHospitalAssignmentRemoved:
		handleErr = r.handleNurseToHospitalAssignmentRemoved(ctx, _event)
	case event.EventTypeMedicalRecordCreated:
		handleErr = r.handleMedicalRecordCreated(ctx, _event)
	case event.EventTypeMedicalRecordDeleted:
		handleErr = r.handleMedicalRecordDeleted(ctx, _event)
	default:
		slog.Error("unknown event type", "event_id", _event.ID, "event_type", _event.Type, "event", outboxEvent.EventJSON)
		return
	}

	if handleErr != nil {
		slog.Error("failed to handle event", "event_id", _event.ID, "event_type", _event.Type, "err", handleErr)
		if err := r.outboxService.UpdateOutboxEventStatus(ctx, outboxEvent.ID, event.OutboxEventStatusFailed); err != nil {
			slog.Error("failed to update event status to failed", "event_id", _event.ID, "event", outboxEvent.EventJSON, "err", err)
		}
		return
	}

	if err := r.outboxService.UpdateOutboxEventStatus(ctx, outboxEvent.ID, event.OutboxEventStatusProcessed); err != nil {
		slog.Error("failed to update event status to processed", "event_id", _event.ID, "event", outboxEvent.EventJSON, "err", err)
	}
}
