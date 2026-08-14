package authorization

import (
	"context"

	"openhealth/internal/domain"
)

type Service struct {
	openfga OpenFGA
}

func NewService(openfga OpenFGA) *Service {
	return &Service{
		openfga: openfga,
	}
}

// AddTuple - writes a generic tuple to OpenFGA.
func (s *Service) AddTuple(ctx context.Context, tuple domain.Tuple) error {
	return s.openfga.AddTuple(ctx, tuple)
}

// RemoveTuple - removes a generic tuple from OpenFGA.
func (s *Service) RemoveTuple(ctx context.Context, tuple domain.Tuple) error {
	return s.openfga.RemoveTuple(ctx, tuple)
}

// Check - asks OpenFGA whether the tuple's user has the requested relation on the object.
func (s *Service) Check(ctx context.Context, tuple domain.Tuple) (bool, error) {
	return s.openfga.Check(ctx, tuple)
}

// AddDoctorAccount - links a doctor to its account.
func (s *Service) AddDoctorAccount(ctx context.Context, doctorID, accountID string) error {
	tuple := domain.NewDoctorAccountTuple(doctorID, accountID)
	return s.openfga.AddTuple(ctx, tuple)
}

// RemoveDoctorAccount - removes the link between a doctor and its account.
func (s *Service) RemoveDoctorAccount(ctx context.Context, doctorID, accountID string) error {
	tuple := domain.NewDoctorAccountTuple(doctorID, accountID)
	return s.openfga.RemoveTuple(ctx, tuple)
}

// AddDoctorToAccountAssignment - assigns a doctor to an account.
func (s *Service) AddDoctorToAccountAssignment(ctx context.Context, accountID, doctorID string) error {
	tuple := domain.NewDoctorToAccountAssignmentTuple(accountID, doctorID)
	return s.openfga.AddTuple(ctx, tuple)
}

// RemoveDoctorToAccountAssignment - removes a doctor to account assignment.
func (s *Service) RemoveDoctorToAccountAssignment(ctx context.Context, accountID, doctorID string) error {
	tuple := domain.NewDoctorToAccountAssignmentTuple(accountID, doctorID)
	return s.openfga.RemoveTuple(ctx, tuple)
}

// AddNurseAccount - links a nurse to its account.
func (s *Service) AddNurseAccount(ctx context.Context, nurseID, accountID string) error {
	tuple := domain.NewNurseAccountTuple(nurseID, accountID)
	return s.openfga.AddTuple(ctx, tuple)
}

// RemoveNurseAccount - removes the link between a nurse and its account.
func (s *Service) RemoveNurseAccount(ctx context.Context, nurseID, accountID string) error {
	tuple := domain.NewNurseAccountTuple(nurseID, accountID)
	return s.openfga.RemoveTuple(ctx, tuple)
}

// AddNurseToHospitalAssignment - assigns a nurse to a hospital.
func (s *Service) AddNurseToHospitalAssignment(ctx context.Context, nurseID, hospitalID string) error {
	tuple := domain.NewNurseToHospitalAssignmentTuple(nurseID, hospitalID)
	return s.openfga.AddTuple(ctx, tuple)
}

// RemoveNurseToHospitalAssignment - removes a nurse to hospital assignment.
func (s *Service) RemoveNurseToHospitalAssignment(ctx context.Context, nurseID, hospitalID string) error {
	tuple := domain.NewNurseToHospitalAssignmentTuple(nurseID, hospitalID)
	return s.openfga.RemoveTuple(ctx, tuple)
}

// AddHospitalToAccountAssignment - assigns a hospital to an account.
func (s *Service) AddHospitalToAccountAssignment(ctx context.Context, accountID, hospitalID string) error {
	tuple := domain.NewHospitalToAccountAssignmentTuple(accountID, hospitalID)
	return s.openfga.AddTuple(ctx, tuple)
}

// RemoveHospitalToAccountAssignment - removes a hospital to account assignment.
func (s *Service) RemoveHospitalToAccountAssignment(ctx context.Context, accountID, hospitalID string) error {
	tuple := domain.NewHospitalToAccountAssignmentTuple(accountID, hospitalID)
	return s.openfga.RemoveTuple(ctx, tuple)
}

// AddMedicalRecordOwner - makes an account the owner of a medical record.
func (s *Service) AddMedicalRecordOwner(ctx context.Context, recordID, accountID string) error {
	tuple := domain.NewMedicalRecordOwnerTuple(recordID, accountID)
	return s.openfga.AddTuple(ctx, tuple)
}

// RemoveMedicalRecordOwner - removes the owner relationship from a medical record.
func (s *Service) RemoveMedicalRecordOwner(ctx context.Context, recordID, accountID string) error {
	tuple := domain.NewMedicalRecordOwnerTuple(recordID, accountID)
	return s.openfga.RemoveTuple(ctx, tuple)
}
