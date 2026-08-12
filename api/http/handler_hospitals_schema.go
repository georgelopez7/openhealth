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
