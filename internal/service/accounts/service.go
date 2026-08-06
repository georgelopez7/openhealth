package accounts

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
		// logger:     logger,
	}
}

// CreateAccount - creates a new account record
func (s *Service) CreateAccount(ctx context.Context, account domain.Account) error {
	err := s.tx.WithTransaction(ctx, func(ctx context.Context) error {
		if err := s.repository.AddAccount(ctx, account); err != nil {
			return err
		}

		// payload := domain.NewAccountCreatedEventPayload(account)
		// event := domain.NewOutboxEvent(domain.EventTypeAccountCreated, payload, domain.OutboxEventStatusPending)

		// if err := s.repository.AddOutboxEvent(ctx, *event); err != nil {
		// 	return err
		// }

		return nil
	})

	return err
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
	return s.repository.UpdateAccount(ctx, account)
}
