package event

import (
	"openhealth/internal/domain"

	"github.com/google/uuid"
)

const (
	EventTypeHospitalCreated EventType = "hospital.created"
	EventTypeHospitalUpdated EventType = "hospital.updated"
	EventTypeHospitalDeleted EventType = "hospital.deleted"

	EventTypeHospitalToAccountAssignmentCreated EventType = "hospital_to_account_assignment.created"
	EventTypeHospitalToAccountAssignmentRemoved EventType = "hospital_to_account_assignment.removed"
)

// hospital.created

type HospitalCreatedEventPayload struct {
	Hospital domain.Hospital `json:"hospital"`
}

func NewHospitalCreatedEvent(hospital domain.Hospital) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeHospitalCreated,
		Payload: HospitalCreatedEventPayload{Hospital: hospital},
	}
}

// hospital.updated

type HospitalUpdatedEventPayload struct {
	Hospital domain.Hospital `json:"hospital"`
}

func NewHospitalUpdatedEvent(hospital domain.Hospital) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeHospitalUpdated,
		Payload: HospitalUpdatedEventPayload{Hospital: hospital},
	}
}

// hospital.deleted

type HospitalDeletedEventPayload struct {
	Hospital domain.Hospital `json:"hospital"`
}

func NewHospitalDeletedEvent(hospital domain.Hospital) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeHospitalDeleted,
		Payload: HospitalDeletedEventPayload{Hospital: hospital},
	}
}

// hospital_to_account_assignment.created

type HospitalToAccountAssignmentCreatedEventPayload struct {
	Assignment domain.HospitalAssignment `json:"assignment"`
}

func NewHospitalToAccountAssignmentCreatedEvent(assignment domain.HospitalAssignment) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeHospitalToAccountAssignmentCreated,
		Payload: HospitalToAccountAssignmentCreatedEventPayload{Assignment: assignment},
	}
}

// hospital_to_account_assignment.removed

type HospitalToAccountAssignmentRemovedEventPayload struct {
	Assignment domain.HospitalAssignment `json:"assignment"`
}

func NewHospitalToAccountAssignmentRemovedEvent(assignment domain.HospitalAssignment) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeHospitalToAccountAssignmentRemoved,
		Payload: HospitalToAccountAssignmentRemovedEventPayload{Assignment: assignment},
	}
}
