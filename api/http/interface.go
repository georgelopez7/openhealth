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
	GetDoctorByID(ctx context.Context, id string) (*domain.Doctor, error)
	AddDoctorToAccountAssignment(ctx context.Context, accountID string, doctorID string) (*domain.DoctorAssignment, error)
}

type NurseSVC interface {
	CreateNurse(ctx context.Context, nurse domain.Nurse) error
	GetNurseByID(ctx context.Context, id string) (*domain.Nurse, error)
	AddNurseToHospitalAssignment(ctx context.Context, nurseID string, hospitalID string) (*domain.NurseToHospitalAssignment, error)
}

type HospitalSVC interface {
	CreateHospital(ctx context.Context, hospital domain.Hospital) error
	GetHospitalByID(ctx context.Context, id string) (*domain.Hospital, error)
	AddHospitalToAccountAssignment(ctx context.Context, accountID string, hospitalID string) (*domain.HospitalAssignment, error)
}

type MedicalRecordSVC interface {
	CreateMedicalRecord(ctx context.Context, record domain.MedicalRecord) error
	GetMedicalRecordByID(ctx context.Context, id string) (*domain.MedicalRecord, error)
	GetMedicalRecordsByAccountID(ctx context.Context, accountID string) ([]domain.MedicalRecord, error)
}
