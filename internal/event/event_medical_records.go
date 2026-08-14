package event

import (
	"openhealth/internal/domain"

	"github.com/google/uuid"
)

const (
	EventTypeMedicalRecordCreated EventType = "medical_record.created"
	EventTypeMedicalRecordDeleted EventType = "medical_record.deleted"
)

// medical_record.created

type MedicalRecordCreatedEventPayload struct {
	Record domain.MedicalRecord `json:"record"`
}

func NewMedicalRecordCreatedEvent(record domain.MedicalRecord) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeMedicalRecordCreated,
		Payload: MedicalRecordCreatedEventPayload{Record: record},
	}
}

// medical_record.deleted

type MedicalRecordDeletedEventPayload struct {
	RecordID  string `json:"record_id"`
	AccountID string `json:"account_id"`
}

func NewMedicalRecordDeletedEvent(recordID, accountID string) Event {
	return Event{
		ID:   uuid.NewString(),
		Type: EventTypeMedicalRecordDeleted,
		Payload: MedicalRecordDeletedEventPayload{
			RecordID:  recordID,
			AccountID: accountID,
		},
	}
}
