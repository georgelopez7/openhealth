package nurses

import (
	"context"
	"openhealth/internal/domain"
)

//go:generate mockgen -source=interface.go -destination=test/mocks.go -package=test

type Repository interface {
	AddNurse(ctx context.Context, nurse domain.Nurse) error
	GetNurseByID(ctx context.Context, id string) (*domain.Nurse, error)
	GetNurses(ctx context.Context, limit int) ([]domain.Nurse, error)
	UpdateNurseStatusByID(ctx context.Context, id string, status domain.NurseStatus) error
	AddNurseToHospitalAssignment(ctx context.Context, assignment domain.NurseToHospitalAssignment) error
	RemoveNurseToHospitalAssignment(ctx context.Context, id string) error
	RemoveAllNurseToHospitalAssignments(ctx context.Context, nurseID string) error
	GetNurseToHospitalAssignmentByID(ctx context.Context, id string) (*domain.NurseToHospitalAssignment, error)
	GetNurseToHospitalAssignments(ctx context.Context, limit int) ([]domain.NurseToHospitalAssignment, error)
	GetNurseToHospitalAssignmentsByNurseID(ctx context.Context, nurseID string) ([]domain.NurseToHospitalAssignment, error)
	GetHospitalIDByNurseID(ctx context.Context, nurseID string) (string, error)
	GetNursesByHospitalID(ctx context.Context, hospitalID string) ([]domain.NurseToHospitalAssignment, error)
}

type TxManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
