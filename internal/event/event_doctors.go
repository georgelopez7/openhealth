package event

import (
	"openhealth/internal/domain"

	"github.com/google/uuid"
)

const (
	EventTypeDoctorCreated EventType = "doctor.created"
	EventTypeDoctorUpdated EventType = "doctor.updated"
	EventTypeDoctorDeleted EventType = "doctor.deleted"

	EventTypeDoctorToAccountAssignmentCreated EventType = "doctor_to_account_assignment.created"
	EventTypeDoctorToAccountAssignmentRemoved EventType = "doctor_to_account_assignment.removed"
)

// doctor.created

type DoctorCreatedEventPayload struct {
	Doctor domain.Doctor `json:"doctor"`
}

func NewDoctorCreatedEvent(doctor domain.Doctor) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeDoctorCreated,
		Payload: DoctorCreatedEventPayload{Doctor: doctor},
	}
}

// doctor.updated

type DoctorUpdatedEventPayload struct {
	Doctor domain.Doctor `json:"doctor"`
}

func NewDoctorUpdatedEvent(doctor domain.Doctor) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeDoctorUpdated,
		Payload: DoctorUpdatedEventPayload{Doctor: doctor},
	}
}

// doctor.deleted

type DoctorDeletedEventPayload struct {
	Doctor domain.Doctor `json:"doctor"`
}

func NewDoctorDeletedEvent(doctor domain.Doctor) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeDoctorDeleted,
		Payload: DoctorDeletedEventPayload{Doctor: doctor},
	}
}

// doctor_to_account_assignment.created

type DoctorToAccountAssignmentCreatedEventPayload struct {
	Assignment domain.DoctorAssignment `json:"assignment"`
}

func NewDoctorToAccountAssignmentCreatedEvent(assignment domain.DoctorAssignment) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeDoctorToAccountAssignmentCreated,
		Payload: DoctorToAccountAssignmentCreatedEventPayload{Assignment: assignment},
	}
}

// doctor_to_account_assignment.removed

type DoctorToAccountAssignmentRemovedEventPayload struct {
	Assignment domain.DoctorAssignment `json:"assignment"`
}

func NewDoctorToAccountAssignmentRemovedEvent(assignment domain.DoctorAssignment) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeDoctorToAccountAssignmentRemoved,
		Payload: DoctorToAccountAssignmentRemovedEventPayload{Assignment: assignment},
	}
}
