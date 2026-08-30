package domain

import (
	"time"

	"github.com/google/uuid"
)

type MedicalRecordStatus string

const (
	MedicalRecordStatusActive   MedicalRecordStatus = "active"
	MedicalRecordStatusArchived MedicalRecordStatus = "archived"
)

type MedicalRecord struct {
	ID          string              `json:"id" db:"id"`
	AccountID   string              `json:"account_id" db:"account_id"`
	Title       string              `json:"title" db:"title"`
	Description string              `json:"description" db:"description"`
	Status      MedicalRecordStatus `json:"status" db:"status"`
	CreatedAt   time.Time           `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at" db:"updated_at"`
}

func NewMedicalRecord(accountID string, title string, description string) *MedicalRecord {
	now := time.Now().UTC()
	return &MedicalRecord{
		ID:          uuid.NewString(),
		AccountID:   accountID,
		Title:       title,
		Description: description,
		Status:      MedicalRecordStatusActive,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
