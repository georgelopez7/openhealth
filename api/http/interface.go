package http

import (
	"context"
	"openhealth/internal/domain"
)

//go:generate mockgen -source=interface.go -destination=test/mocks.go -package=test

type AccountSVC interface {
	CreateAccount(ctx context.Context, account domain.Account) error
	GetAccountByID(ctx context.Context, id string) (*domain.Account, error)
	GetAccounts(ctx context.Context, limit int) ([]domain.Account, error)
	UpdateAccount(ctx context.Context, account domain.Account) error
	ArchiveAccount(ctx context.Context, id string) error
}

type DoctorSVC interface {
	CreateDoctor(ctx context.Context, doctor domain.Doctor) error
	AddDoctorToAccountAssignment(ctx context.Context, accountID string, doctorID string) (*domain.DoctorAssignment, error)
}
