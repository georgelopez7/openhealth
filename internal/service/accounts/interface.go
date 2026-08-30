package accounts

import (
	"context"
	"openhealth/internal/domain"
	"openhealth/internal/event"
)

//go:generate mockgen -source=interface.go -destination=test/mocks.go -package=test

type Repository interface {
	AddAccount(ctx context.Context, account domain.Account) error
	GetAccountByID(ctx context.Context, id string) (*domain.Account, error)
	GetAccounts(ctx context.Context, limit int) ([]domain.Account, error)
	UpdateAccount(ctx context.Context, account domain.Account) error
	ArchiveAccount(ctx context.Context, id string) error
	GetDoctorIDByAccountID(ctx context.Context, accountID string) (string, error)
	GetHospitalIDByAccountID(ctx context.Context, accountID string) (string, error)
	AddOutboxEvent(ctx context.Context, outboxEvent event.OutboxEvent) error
}

type TxManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
