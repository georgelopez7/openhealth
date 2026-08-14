package hospitals

import (
	"context"
	"openhealth/internal/domain"
	"openhealth/internal/event"
)

//go:generate mockgen -source=interface.go -destination=test/mocks.go -package=test

type Repository interface {
	AddHospital(ctx context.Context, hospital domain.Hospital) error
	GetHospitalByID(ctx context.Context, id string) (*domain.Hospital, error)
	GetHospitals(ctx context.Context, limit int) ([]domain.Hospital, error)
	UpdateHospital(ctx context.Context, hospital domain.Hospital) error
	ArchiveHospital(ctx context.Context, id string) error
	ResetHospitals(ctx context.Context) error
	AddHospitalAssignment(ctx context.Context, assignment domain.HospitalAssignment) error
	RemoveHospitalToAccountAssignment(ctx context.Context, id string) error
	GetHospitalAssignmentByID(ctx context.Context, id string) (*domain.HospitalAssignment, error)
	GetHospitalAssignments(ctx context.Context, limit int) ([]domain.HospitalAssignment, error)
	GetHospitalIDByAccountID(ctx context.Context, accountID string) (string, error)
	AddOutboxEvent(ctx context.Context, outboxEvent event.OutboxEvent) error
}

type TxManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
