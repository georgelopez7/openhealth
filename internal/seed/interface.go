package seed

import (
	"context"

	"openhealth/internal/domain"
)

//go:generate mockgen -source=interface.go -destination=test/mocks.go -package=test

type AccountService interface {
	CreateAccount(ctx context.Context, account domain.Account) error
	GetAccounts(ctx context.Context, limit int) ([]domain.Account, error)
}

type DoctorService interface {
	CreateDoctor(ctx context.Context, doctor domain.Doctor) error
	AddDoctorToAccountAssignment(ctx context.Context, accountID, doctorID string) (*domain.DoctorAssignment, error)
}

type NurseService interface {
	CreateNurse(ctx context.Context, nurse domain.Nurse) error
	AddNurseToHospitalAssignment(ctx context.Context, nurseID, hospitalID string) (*domain.NurseToHospitalAssignment, error)
}

type HospitalService interface {
	CreateHospital(ctx context.Context, hospital domain.Hospital) error
	AddHospitalToAccountAssignment(ctx context.Context, accountID, hospitalID string) (*domain.HospitalAssignment, error)
}

type MedicalRecordService interface {
	CreateMedicalRecord(ctx context.Context, record domain.MedicalRecord) error
}
