package doctors

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

// CreateDoctor - creates a new doctor record
func (s *Service) CreateDoctor(ctx context.Context, doctor domain.Doctor) error {
	return s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		return s.repository.AddDoctor(ctx, doctor)
	})
}

// GetDoctorByID - get a doctor by its ID
func (s *Service) GetDoctorByID(ctx context.Context, id string) (*domain.Doctor, error) {
	doctor, err := s.repository.GetDoctorByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if doctor == nil {
		return nil, domain.DoctorNotFoundError
	}

	return doctor, nil
}

// GetDoctors - get doctors up to the provided limit
func (s *Service) GetDoctors(ctx context.Context, limit int) ([]domain.Doctor, error) {
	doctors, err := s.repository.GetDoctors(ctx, limit)
	if err != nil {
		return nil, err
	}

	return doctors, nil
}

// UpdateDoctorStatusByID - updates the status of a doctor; expires open assignments when archiving
func (s *Service) UpdateDoctorStatusByID(ctx context.Context, id string, status domain.DoctorStatus) error {
	return s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		if err := s.repository.UpdateDoctorStatusByID(ctx, id, status); err != nil {
			return err
		}

		if status == domain.DoctorStatusArchived {
			return s.repository.ExpireDoctorAssignments(ctx, id)
		}

		return nil
	})
}

// AssignDoctorToAccount - assigns a doctor to an account
func (s *Service) AssignDoctorToAccount(ctx context.Context, accountID string, doctorID string) error {
	return s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		return s.repository.AssignDoctorToAccount(ctx, accountID, doctorID)
	})
}
