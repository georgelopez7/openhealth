package doctors

import (
	"context"
	"openhealth/internal/domain"
)

//go:generate mockgen -source=interface.go -destination=test/mocks.go -package=test

type Repository interface {
	AddDoctor(ctx context.Context, doctor domain.Doctor) error
	GetDoctorByID(ctx context.Context, id string) (*domain.Doctor, error)
	GetDoctorByAccountID(ctx context.Context, accountID string) (*domain.Doctor, error)
	GetDoctors(ctx context.Context, limit int) ([]domain.Doctor, error)
	UpdateDoctorStatusByID(ctx context.Context, id string, status domain.DoctorStatus) error
	AddDoctorToAccountAssignment(ctx context.Context, assignment domain.DoctorAssignment) error
	RemoveDoctorAssignment(ctx context.Context, id string) error
	RemoveAllDoctorToAccountAssignments(ctx context.Context, doctorID string) error
	GetDoctorAssignmentByID(ctx context.Context, id string) (*domain.DoctorAssignment, error)
	GetDoctorAssignments(ctx context.Context, limit int) ([]domain.DoctorAssignment, error)
	GetDoctorAssignmentsByDoctorID(ctx context.Context, doctorID string) ([]domain.DoctorAssignment, error)
	GetDoctorIDByAccountID(ctx context.Context, accountID string) (string, error)
}

type TxManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
