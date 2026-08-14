package doctors

import (
	"context"

	"openhealth/internal/domain"
	"openhealth/internal/event"
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

func (s *Service) addOutboxEvent(ctx context.Context, evt event.Event) error {
	outboxEvent, err := event.NewOutboxEvent(evt)
	if err != nil {
		return err
	}

	return s.repository.AddOutboxEvent(ctx, *outboxEvent)
}

// CreateDoctor - creates a new doctor record
func (s *Service) CreateDoctor(ctx context.Context, doctor domain.Doctor) error {
	return s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		if err := s.repository.AddDoctor(ctx, doctor); err != nil {
			return err
		}

		return s.addOutboxEvent(ctx, event.NewDoctorCreatedEvent(doctor))
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

// GetDoctorByAccountID - get a doctor by its associated account ID
func (s *Service) GetDoctorByAccountID(ctx context.Context, accountID string) (*domain.Doctor, error) {
	doctor, err := s.repository.GetDoctorByAccountID(ctx, accountID)
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

// UpdateDoctorStatusByID - updates the status of a doctor by ID
func (s *Service) UpdateDoctorStatusByID(ctx context.Context, id string, status domain.DoctorStatus) error {
	return s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		if err := s.repository.UpdateDoctorStatusByID(ctx, id, status); err != nil {
			return err
		}

		if status != domain.DoctorStatusArchived {
			return s.addOutboxEvent(ctx, event.NewDoctorUpdatedEvent(domain.Doctor{ID: id, Status: status}))
		}

		assignments, err := s.repository.GetDoctorAssignmentsByDoctorID(ctx, id)
		if err != nil {
			return err
		}

		if err := s.repository.RemoveAllDoctorToAccountAssignments(ctx, id); err != nil {
			return err
		}

		for _, assignment := range assignments {
			if err := s.addOutboxEvent(ctx, event.NewDoctorToAccountAssignmentRemovedEvent(assignment)); err != nil {
				return err
			}
		}

		return s.addOutboxEvent(ctx, event.NewDoctorUpdatedEvent(domain.Doctor{ID: id, Status: status}))
	})
}

// AddDoctorToAccountAssignment - assigns a doctor to an account
func (s *Service) AddDoctorToAccountAssignment(ctx context.Context, accountID string, doctorID string) (*domain.DoctorAssignment, error) {
	assignment := domain.NewDoctorToAccountAssignment(accountID, doctorID)

	err := s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		if err := s.repository.AddDoctorToAccountAssignment(ctx, *assignment); err != nil {
			return err
		}

		return s.addOutboxEvent(ctx, event.NewDoctorToAccountAssignmentCreatedEvent(*assignment))
	})
	if err != nil {
		return nil, err
	}

	return assignment, nil
}
