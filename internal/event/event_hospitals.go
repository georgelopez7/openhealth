package event

import (
	"openhealth/internal/domain"

	"github.com/google/uuid"
)

const (
	EventTypeHospitalCreated EventType = "hospital.created"
	EventTypeHospitalUpdated EventType = "hospital.updated"
	EventTypeHospitalDeleted EventType = "hospital.deleted"
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
