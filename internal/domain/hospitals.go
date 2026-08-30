package domain

import (
	"time"

	"github.com/google/uuid"
)

type HospitalStatus string

const (
	HospitalStatusActive   HospitalStatus = "active"
	HospitalStatusArchived HospitalStatus = "archived"
)

type Hospital struct {
	ID        string         `json:"id" db:"id"`
	Name      string         `json:"name" db:"name"`
	Status    HospitalStatus `json:"status" db:"status"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
}

func NewHospital(name string) *Hospital {
	return &Hospital{
		ID:        uuid.NewString(),
		Name:      name,
		Status:    HospitalStatusActive,
		CreatedAt: time.Now().UTC(),
	}
}

type HospitalAssignment struct {
	ID         string     `json:"id" db:"id"`
	AccountID  string     `json:"account_id" db:"account_id"`
	HospitalID string     `json:"hospital_id" db:"hospital_id"`
	ValidFrom  time.Time  `json:"valid_from" db:"valid_from"`
	ValidTo    *time.Time `json:"valid_to" db:"valid_to"`
}

func NewHospitalAssignment(accountID string, hospitalID string) *HospitalAssignment {
	return &HospitalAssignment{
		ID:         uuid.NewString(),
		AccountID:  accountID,
		HospitalID: hospitalID,
		ValidFrom:  time.Now().UTC(),
	}
}
