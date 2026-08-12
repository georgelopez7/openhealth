package medicalrecords

import (
	"context"

	"openhealth/internal/domain"
)

type Service struct {
	tx         TxManager
	repository Repository
}

func NewService(tx TxManager, repository Repository) *Service {
	return &Service{
		tx:         tx,
		repository: repository,
	}
}

// CreateMedicalRecord - creates a new medical record
func (s *Service) CreateMedicalRecord(ctx context.Context, record domain.MedicalRecord) error {
	return s.repository.AddMedicalRecord(ctx, record)
}

// GetMedicalRecordByID - get a medical record by its ID
func (s *Service) GetMedicalRecordByID(ctx context.Context, id string) (*domain.MedicalRecord, error) {
	record, err := s.repository.GetMedicalRecordByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if record == nil {
		return nil, domain.MedicalRecordNotFoundError
	}

	return record, nil
}

// UpdateMedicalRecord - updates an existing medical record
func (s *Service) UpdateMedicalRecord(ctx context.Context, record domain.MedicalRecord) error {
	return s.repository.UpdateMedicalRecord(ctx, record)
}

// ArchiveMedicalRecord - archives an existing medical record
func (s *Service) ArchiveMedicalRecord(ctx context.Context, id string) error {
	return s.repository.ArchiveMedicalRecord(ctx, id)
}
