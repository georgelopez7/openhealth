package event

import (
	"openhealth/internal/domain"

	"github.com/google/uuid"
)

const (
	EventTypeNurseCreated EventType = "nurse.created"
	EventTypeNurseUpdated EventType = "nurse.updated"
	EventTypeNurseDeleted EventType = "nurse.deleted"
)

// nurse.created

type NurseCreatedEventPayload struct {
	Nurse domain.Nurse `json:"nurse"`
}

func NewNurseCreatedEvent(nurse domain.Nurse) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeNurseCreated,
		Payload: NurseCreatedEventPayload{Nurse: nurse},
	}
}

// nurse.updated

type NurseUpdatedEventPayload struct {
	Nurse domain.Nurse `json:"nurse"`
}

func NewNurseUpdatedEvent(nurse domain.Nurse) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeNurseUpdated,
		Payload: NurseUpdatedEventPayload{Nurse: nurse},
	}
}

// nurse.deleted

type NurseDeletedEventPayload struct {
	Nurse domain.Nurse `json:"nurse"`
}

func NewNurseDeletedEvent(nurse domain.Nurse) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeNurseDeleted,
		Payload: NurseDeletedEventPayload{Nurse: nurse},
	}
}
