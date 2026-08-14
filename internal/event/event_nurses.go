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

const (
	EventTypeNurseToHospitalAssignmentCreated EventType = "nurse_to_hospital_assignment.created"
	EventTypeNurseToHospitalAssignmentRemoved EventType = "nurse_to_hospital_assignment.removed"
)

// nurse_to_hospital_assignment.created

type NurseToHospitalAssignmentCreatedEventPayload struct {
	Assignment domain.NurseToHospitalAssignment `json:"assignment"`
}

func NewNurseToHospitalAssignmentCreatedEvent(assignment domain.NurseToHospitalAssignment) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeNurseToHospitalAssignmentCreated,
		Payload: NurseToHospitalAssignmentCreatedEventPayload{Assignment: assignment},
	}
}

// nurse_to_hospital_assignment.removed

type NurseToHospitalAssignmentRemovedEventPayload struct {
	Assignment domain.NurseToHospitalAssignment `json:"assignment"`
}

func NewNurseToHospitalAssignmentRemovedEvent(assignment domain.NurseToHospitalAssignment) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeNurseToHospitalAssignmentRemoved,
		Payload: NurseToHospitalAssignmentRemovedEventPayload{Assignment: assignment},
	}
}
