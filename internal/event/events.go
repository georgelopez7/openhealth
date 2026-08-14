package event

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type EventType string

type Event struct {
	ID        string    `json:"id"`
	Type      EventType `json:"type"`
	Payload   any       `json:"payload"`
	CreatedAt string    `json:"createdAt"`
}

type OutboxEventStatus string

const (
	OutboxEventStatusPending   OutboxEventStatus = "pending"
	OutboxEventStatusProcessed OutboxEventStatus = "processed"
	OutboxEventStatusFailed    OutboxEventStatus = "failed"
)

type OutboxEvent struct {
	ID        string            `json:"id" db:"id"`
	Status    OutboxEventStatus `json:"status" db:"status"`
	EventID   string            `json:"event_id" db:"event_id"`
	EventType EventType         `json:"event_type" db:"event_type"`
	EventJSON json.RawMessage   `json:"event_json" db:"event_json"`
	CreatedAt time.Time         `json:"created_at" db:"created_at"`
}

func NewOutboxEvent(event Event) (*OutboxEvent, error) {
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	return &OutboxEvent{
		ID:        uuid.NewString(),
		Status:    OutboxEventStatusPending,
		EventID:   event.ID,
		EventType: event.Type,
		EventJSON: eventJSON,
		CreatedAt: time.Now().UTC(),
	}, nil
}
