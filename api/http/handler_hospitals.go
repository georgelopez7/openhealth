package http

import (
	"context"

	"openhealth/internal/domain"
)

func (s *Server) CreateHospitalHandler(ctx context.Context, input *CreateHospitalInput) (*CreateHospitalResponse, error) {
	hospital := domain.NewHospital(input.Body.Name)

	err := s.HospitalSVC.CreateHospital(ctx, *hospital)
	if err != nil {
		return nil, err
	}

	resp := &CreateHospitalResponse{}
	resp.Body.Hospital = *hospital

	return resp, nil
}
