package http

import "openhealth/internal/domain"

type CreateHospitalInput struct {
	Body struct {
		Name string `json:"name"`
	}
}

type CreateHospitalResponse struct {
	Body struct {
		Hospital domain.Hospital `json:"hospital"`
	}
}

type GetHospitalByIDInput struct {
	HospitalID string `path:"hospitalID"`
}

type GetHospitalByIDResponse struct {
	Body struct {
		Hospital domain.Hospital `json:"hospital"`
	}
}

type CreateHospitalToAccountAssignmentInput struct {
	Body struct {
		AccountID  string `json:"account_id"`
		HospitalID string `json:"hospital_id"`
	}
}

type CreateHospitalToAccountAssignmentResponse struct {
	Body struct {
		Assignment domain.HospitalAssignment `json:"assignment"`
	}
}
