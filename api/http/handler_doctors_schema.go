package http

import "openhealth/internal/domain"

type CreateDoctorInput struct {
	Body struct {
		AccountID string `json:"account_id"`
	}
}

type CreateDoctorResponse struct {
	Body struct {
		Doctor domain.Doctor `json:"doctor"`
	}
}

type GetDoctorByIDInput struct {
	DoctorID string `path:"doctorID"`
}

type GetDoctorByIDResponse struct {
	Body struct {
		Doctor domain.Doctor `json:"doctor"`
	}
}

type CreateDoctorToAccountAssignmentInput struct {
	Body struct {
		AccountID string `json:"account_id"`
		DoctorID  string `json:"doctor_id"`
	}
}

type CreateDoctorToAccountAssignmentResponse struct {
	Body struct {
		Assignment domain.DoctorAssignment `json:"assignment"`
	}
}
