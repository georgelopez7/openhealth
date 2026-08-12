package doctors

import (
	"context"
	"openhealth/internal/domain"
)

//go:generate mockgen -source=interface.go -destination=test/mocks.go -package=test

type Repository interface {
	AddDoctor(ctx context.Context, doctor domain.Doctor) error
	GetDoctorByID(ctx context.Context, id string) (*domain.Doctor, error)
	GetDoctors(ctx context.Context, limit int) ([]domain.Doctor, error)
	UpdateDoctorStatusByID(ctx context.Context, id string, status domain.DoctorStatus) error
	AssignDoctorToAccount(ctx context.Context, accountID string, doctorID string) error
	ExpireDoctorAssignments(ctx context.Context, doctorID string) error
}

type TxManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
