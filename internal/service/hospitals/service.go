package hospitals

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

// CreateHospital - creates a new hospital record
func (s *Service) CreateHospital(ctx context.Context, hospital domain.Hospital) error {
	return s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		if err := s.repository.AddHospital(ctx, hospital); err != nil {
			return err
		}

		return s.addOutboxEvent(ctx, event.NewHospitalCreatedEvent(hospital))
	})
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
	return s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		if err := s.repository.UpdateHospital(ctx, hospital); err != nil {
			return err
		}

		return s.addOutboxEvent(ctx, event.NewHospitalUpdatedEvent(hospital))
	})
}

// ArchiveHospital - archives an existing hospital
func (s *Service) ArchiveHospital(ctx context.Context, id string) error {
	return s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		hospital, err := s.repository.GetHospitalByID(ctx, id)
		if err != nil {
			return err
		}

		if hospital == nil {
			return domain.HospitalNotFoundError
		}

		if err := s.repository.ArchiveHospital(ctx, id); err != nil {
			return err
		}

		return s.addOutboxEvent(ctx, event.NewHospitalDeletedEvent(*hospital))
	})
}

// AddHospitalToAccountAssignment - assigns a hospital to an account
func (s *Service) AddHospitalToAccountAssignment(ctx context.Context, accountID string, hospitalID string) (*domain.HospitalAssignment, error) {
	assignment := domain.NewHospitalAssignment(accountID, hospitalID)

	err := s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		if err := s.repository.AddHospitalAssignment(ctx, *assignment); err != nil {
			return err
		}

		return s.addOutboxEvent(ctx, event.NewHospitalToAccountAssignmentCreatedEvent(*assignment))
	})
	if err != nil {
		return nil, err
	}

	return assignment, nil
}

// RemoveHospitalToAccountAssignment - removes a hospital to account assignment by ID
func (s *Service) RemoveHospitalToAccountAssignment(ctx context.Context, id string) error {
	return s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		assignment, err := s.repository.GetHospitalAssignmentByID(ctx, id)
		if err != nil {
			return err
		}

		if assignment == nil {
			return nil
		}

		if err := s.repository.RemoveHospitalToAccountAssignment(ctx, id); err != nil {
			return err
		}

		return s.addOutboxEvent(ctx, event.NewHospitalToAccountAssignmentRemovedEvent(*assignment))
	})
}
