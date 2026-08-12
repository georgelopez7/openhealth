package http

import "openhealth/internal/domain"

type CreateNurseInput struct {
	Body struct {
		AccountID string `json:"account_id"`
	}
}

type CreateNurseResponse struct {
	Body struct {
		Nurse domain.Nurse `json:"nurse"`
	}
}

type GetNurseByIDInput struct {
	NurseID string `path:"nurseID"`
}

type GetNurseByIDResponse struct {
	Body struct {
		Nurse domain.Nurse `json:"nurse"`
	}
}

type CreateNurseToHospitalAssignmentInput struct {
	Body struct {
		NurseID    string `json:"nurse_id"`
		HospitalID string `json:"hospital_id"`
	}
}

type CreateNurseToHospitalAssignmentResponse struct {
	Body struct {
		Assignment domain.NurseToHospitalAssignment `json:"assignment"`
	}
}
