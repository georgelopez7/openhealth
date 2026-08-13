package medicalrecords

import (
	"context"

	"openhealth/internal/domain"
)

//go:generate mockgen -source=interface.go -destination=test/mocks.go -package=test

type Repository interface {
	AddMedicalRecord(ctx context.Context, record domain.MedicalRecord) error
	GetMedicalRecordByID(ctx context.Context, id string) (*domain.MedicalRecord, error)
	GetMedicalRecords(ctx context.Context) ([]domain.MedicalRecord, error)
	GetMedicalRecordsByAccountID(ctx context.Context, accountID string) ([]domain.MedicalRecord, error)
	UpdateMedicalRecord(ctx context.Context, record domain.MedicalRecord) error
	ArchiveMedicalRecord(ctx context.Context, id string) error
}

type TxManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
