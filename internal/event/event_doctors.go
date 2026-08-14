package event

import (
	"openhealth/internal/domain"

	"github.com/google/uuid"
)

const (
	EventTypeDoctorCreated EventType = "doctor.created"
	EventTypeDoctorUpdated EventType = "doctor.updated"
	EventTypeDoctorDeleted EventType = "doctor.deleted"
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
