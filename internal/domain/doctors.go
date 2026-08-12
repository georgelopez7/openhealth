package domain

import (
	"time"

	"github.com/google/uuid"
)

type DoctorStatus string

const (
	DoctorStatusActive   DoctorStatus = "active"
	DoctorStatusArchived DoctorStatus = "archived"
)

type Doctor struct {
	ID        string       `json:"id" db:"id"`
	AccountID string       `json:"account_id" db:"account_id"`
	Status    DoctorStatus `json:"status" db:"status"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
}

func NewDoctor(accountID string) *Doctor {
	return &Doctor{
		ID:        uuid.NewString(),
		AccountID: accountID,
		Status:    DoctorStatusActive,
		CreatedAt: time.Now().UTC(),
	}
}
