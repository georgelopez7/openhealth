package accounts

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

// CreateAccount - creates a new account record
func (s *Service) CreateAccount(ctx context.Context, account domain.Account) error {
	return s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		if err := s.repository.AddAccount(ctx, account); err != nil {
			return err
		}

		return s.addOutboxEvent(ctx, event.NewAccountCreatedEvent(account))
	})
}

// GetAccountByID - get an account by its ID
func (s *Service) GetAccountByID(ctx context.Context, id string) (*domain.Account, error) {
	account, err := s.repository.GetAccountByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, domain.AccountNotFoundError
	}

	doctorID, err := s.repository.GetDoctorIDByAccountID(ctx, id)
	if err != nil {
		return nil, err
	}

	account.DoctorID = doctorID

	hospitalID, err := s.repository.GetHospitalIDByAccountID(ctx, id)
	if err != nil {
		return nil, err
	}

	account.HospitalID = hospitalID

	return account, nil
}

// GetAccounts - get accounts up to the provided limit
func (s *Service) GetAccounts(ctx context.Context, limit int) ([]domain.Account, error) {
	accounts, err := s.repository.GetAccounts(ctx, limit)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// UpdateAccount - updates an existing account record
func (s *Service) UpdateAccount(ctx context.Context, account domain.Account) error {
	return s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		if err := s.repository.UpdateAccount(ctx, account); err != nil {
			return err
		}

		return s.addOutboxEvent(ctx, event.NewAccountUpdatedEvent(account))
	})
}

// ArchiveAccount - archives an existing account
func (s *Service) ArchiveAccount(ctx context.Context, id string) error {
	return s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		account, err := s.repository.GetAccountByID(ctx, id)
		if err != nil {
			return err
		}

		if account == nil {
			return domain.AccountNotFoundError
		}

		if err := s.repository.ArchiveAccount(ctx, id); err != nil {
			return err
		}

		return s.addOutboxEvent(ctx, event.NewAccountDeletedEvent(*account))
	})
}
