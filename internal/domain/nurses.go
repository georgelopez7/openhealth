package domain

import (
	"time"

	"github.com/google/uuid"
)

type NurseStatus string

const (
	NurseStatusActive   NurseStatus = "active"
	NurseStatusArchived NurseStatus = "archived"
)

type Nurse struct {
	ID        string      `json:"id" db:"id"`
	AccountID string      `json:"account_id" db:"account_id"`
	Status    NurseStatus `json:"status" db:"status"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
}

func NewNurse(accountID string) *Nurse {
	return &Nurse{
		ID:        uuid.NewString(),
		AccountID: accountID,
		Status:    NurseStatusActive,
		CreatedAt: time.Now().UTC(),
	}
}

type NurseToHospitalAssignment struct {
	ID         string     `json:"id" db:"id"`
	NurseID    string     `json:"nurse_id" db:"nurse_id"`
	HospitalID string     `json:"hospital_id" db:"hospital_id"`
	ValidFrom  time.Time  `json:"valid_from" db:"valid_from"`
	ValidTo    *time.Time `json:"valid_to" db:"valid_to"`
}

func NewNurseToHospitalAssignment(nurseID string, hospitalID string) *NurseToHospitalAssignment {
	return &NurseToHospitalAssignment{
		ID:         uuid.NewString(),
		NurseID:    nurseID,
		HospitalID: hospitalID,
		ValidFrom:  time.Now().UTC(),
	}
}
