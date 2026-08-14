package medicalrecords

import (
	"context"

	"openhealth/internal/domain"
)

type Service struct {
	tx               TxManager
	repository       Repository
	authorizationSVC AuthorizationSVC
}

func NewService(tx TxManager, repository Repository, authorizationSVC AuthorizationSVC) *Service {
	return &Service{
		tx:               tx,
		repository:       repository,
		authorizationSVC: authorizationSVC,
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

// GetMedicalRecordAccess - checks whether an account can view and/or edit a medical record
func (s *Service) GetMedicalRecordAccess(ctx context.Context, recordID, accountID string) (canView bool, canEdit bool, err error) {
	record, err := s.repository.GetMedicalRecordByID(ctx, recordID)
	if err != nil {
		return false, false, err
	}

	if record == nil {
		return false, false, domain.MedicalRecordNotFoundError
	}

	canView, err = s.authorizationSVC.Check(ctx, domain.NewMedicalRecordCanViewTuple(recordID, accountID))
	if err != nil {
		return false, false, err
	}

	canEdit, err = s.authorizationSVC.Check(ctx, domain.NewMedicalRecordCanEditTuple(recordID, accountID))
	if err != nil {
		return false, false, err
	}

	return canView, canEdit, nil
}

// GetMedicalRecords - get all medical records
func (s *Service) GetMedicalRecords(ctx context.Context) ([]domain.MedicalRecord, error) {
	records, err := s.repository.GetMedicalRecords(ctx)
	if err != nil {
		return nil, err
	}

	return records, nil
}

// GetMedicalRecordsByAccountID - get medical records for an account
func (s *Service) GetMedicalRecordsByAccountID(ctx context.Context, accountID string) ([]domain.MedicalRecord, error) {
	records, err := s.repository.GetMedicalRecordsByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	return records, nil
}

// UpdateMedicalRecord - updates an existing medical record
func (s *Service) UpdateMedicalRecord(ctx context.Context, record domain.MedicalRecord) error {
	return s.repository.UpdateMedicalRecord(ctx, record)
}

// ArchiveMedicalRecord - archives an existing medical record
func (s *Service) ArchiveMedicalRecord(ctx context.Context, id string) error {
	return s.repository.ArchiveMedicalRecord(ctx, id)
}
