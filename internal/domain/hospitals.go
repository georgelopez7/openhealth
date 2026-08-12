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
