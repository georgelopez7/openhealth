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

type DoctorAssignment struct {
	ID        string     `json:"id" db:"id"`
	AccountID string     `json:"account_id" db:"account_id"`
	DoctorID  string     `json:"doctor_id" db:"doctor_id"`
	ValidFrom time.Time  `json:"valid_from" db:"valid_from"`
	ValidTo   *time.Time `json:"valid_to" db:"valid_to"`
}

func NewDoctorToAccountAssignment(accountID string, doctorID string) *DoctorAssignment {
	return &DoctorAssignment{
		ID:        uuid.NewString(),
		AccountID: accountID,
		DoctorID:  doctorID,
		ValidFrom: time.Now().UTC(),
	}
}
