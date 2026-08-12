package hospitals

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

// CreateHospital - creates a new hospital record
func (s *Service) CreateHospital(ctx context.Context, hospital domain.Hospital) error {
	return s.repository.AddHospital(ctx, hospital)
}

// GetHospitalByID - get a hospital by its ID
func (s *Service) GetHospitalByID(ctx context.Context, id string) (*domain.Hospital, error) {
	hospital, err := s.repository.GetHospitalByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if hospital == nil {
		return nil, domain.HospitalNotFoundError
	}

	return hospital, nil
}

// GetHospitals - get hospitals up to the provided limit
func (s *Service) GetHospitals(ctx context.Context, limit int) ([]domain.Hospital, error) {
	hospitals, err := s.repository.GetHospitals(ctx, limit)
	if err != nil {
		return nil, err
	}

	return hospitals, nil
}

// UpdateHospital - updates an existing hospital
func (s *Service) UpdateHospital(ctx context.Context, hospital domain.Hospital) error {
	return s.repository.UpdateHospital(ctx, hospital)
}

// ArchiveHospital - archives an existing hospital
func (s *Service) ArchiveHospital(ctx context.Context, id string) error {
	return s.repository.ArchiveHospital(ctx, id)
}

// AddHospitalToAccountAssignment - assigns a hospital to an account
func (s *Service) AddHospitalToAccountAssignment(ctx context.Context, accountID string, hospitalID string) error {
	assignment := domain.NewHospitalAssignment(accountID, hospitalID)

	return s.repository.AddHospitalAssignment(ctx, *assignment)
}

// RemoveHospitalToAccountAssignment - removes a hospital to account assignment by ID
func (s *Service) RemoveHospitalToAccountAssignment(ctx context.Context, id string) error {
	return s.repository.RemoveHospitalToAccountAssignment(ctx, id)
}
