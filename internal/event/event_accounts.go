package event

import (
	"openhealth/internal/domain"

	"github.com/google/uuid"
)

const (
	EventTypeAccountCreated EventType = "account.created"
	EventTypeAccountUpdated EventType = "account.updated"
	EventTypeAccountDeleted EventType = "account.deleted"
)

// account.created

type AccountCreatedEventPayload struct {
	Account domain.Account `json:"account"`
}

func NewAccountCreatedEvent(account domain.Account) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeAccountCreated,
		Payload: AccountCreatedEventPayload{Account: account},
	}
}

// account.updated

type AccountUpdatedEventPayload struct {
	Account domain.Account `json:"account"`
}

func NewAccountUpdatedEvent(account domain.Account) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeAccountUpdated,
		Payload: AccountUpdatedEventPayload{Account: account},
	}
}

// account.deleted

type AccountDeletedEventPayload struct {
	Account domain.Account `json:"account"`
}

func NewAccountDeletedEvent(account domain.Account) Event {
	return Event{
		ID:      uuid.NewString(),
		Type:    EventTypeAccountDeleted,
		Payload: AccountDeletedEventPayload{Account: account},
	}
}
