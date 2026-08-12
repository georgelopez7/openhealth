package nurses

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

// CreateNurse - creates a new nurse record
func (s *Service) CreateNurse(ctx context.Context, nurse domain.Nurse) error {
	return s.repository.AddNurse(ctx, nurse)
}

// GetNurseByID - get a nurse by its ID
func (s *Service) GetNurseByID(ctx context.Context, id string) (*domain.Nurse, error) {
	nurse, err := s.repository.GetNurseByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if nurse == nil {
		return nil, domain.NurseNotFoundError
	}

	return nurse, nil
}

// GetNurses - get nurses up to the provided limit
func (s *Service) GetNurses(ctx context.Context, limit int) ([]domain.Nurse, error) {
	nurses, err := s.repository.GetNurses(ctx, limit)
	if err != nil {
		return nil, err
	}

	return nurses, nil
}

// UpdateNurseStatusByID - updates the status of a nurse by ID
func (s *Service) UpdateNurseStatusByID(ctx context.Context, id string, status domain.NurseStatus) error {
	return s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		if err := s.repository.UpdateNurseStatusByID(ctx, id, status); err != nil {
			return err
		}

		if status != domain.NurseStatusArchived {
			return nil
		}

		return s.repository.RemoveAllNurseToHospitalAssignments(ctx, id)
	})
}

// AddNurseToHospitalAssignment - assigns a nurse to a hospital
func (s *Service) AddNurseToHospitalAssignment(ctx context.Context, nurseID string, hospitalID string) error {
	assignment := domain.NewNurseToHospitalAssignment(nurseID, hospitalID)

	return s.repository.AddNurseToHospitalAssignment(ctx, *assignment)
}
