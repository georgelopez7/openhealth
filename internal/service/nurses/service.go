package nurses

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

// CreateNurse - creates a new nurse record
func (s *Service) CreateNurse(ctx context.Context, nurse domain.Nurse) error {
	return s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		if err := s.repository.AddNurse(ctx, nurse); err != nil {
			return err
		}

		return s.addOutboxEvent(ctx, event.NewNurseCreatedEvent(nurse))
	})
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

// GetNurseByAccountID - get a nurse by its associated account ID
func (s *Service) GetNurseByAccountID(ctx context.Context, accountID string) (*domain.Nurse, error) {
	nurse, err := s.repository.GetNurseByAccountID(ctx, accountID)
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
			return s.addOutboxEvent(ctx, event.NewNurseUpdatedEvent(domain.Nurse{ID: id, Status: status}))
		}

		assignments, err := s.repository.GetNurseToHospitalAssignmentsByNurseID(ctx, id)
		if err != nil {
			return err
		}

		if err := s.repository.RemoveAllNurseToHospitalAssignments(ctx, id); err != nil {
			return err
		}

		for _, assignment := range assignments {
			if err := s.addOutboxEvent(ctx, event.NewNurseToHospitalAssignmentRemovedEvent(assignment)); err != nil {
				return err
			}
		}

		return s.addOutboxEvent(ctx, event.NewNurseUpdatedEvent(domain.Nurse{ID: id, Status: status}))
	})
}

// AddNurseToHospitalAssignment - assigns a nurse to a hospital
func (s *Service) AddNurseToHospitalAssignment(ctx context.Context, nurseID string, hospitalID string) (*domain.NurseToHospitalAssignment, error) {
	assignment := domain.NewNurseToHospitalAssignment(nurseID, hospitalID)

	err := s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		if err := s.repository.AddNurseToHospitalAssignment(ctx, *assignment); err != nil {
			return err
		}

		return s.addOutboxEvent(ctx, event.NewNurseToHospitalAssignmentCreatedEvent(*assignment))
	})
	if err != nil {
		return nil, err
	}

	return assignment, nil
}
