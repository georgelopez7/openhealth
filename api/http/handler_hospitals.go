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

func (s *Server) CreateHospitalToAccountAssignmentHandler(ctx context.Context, input *CreateHospitalToAccountAssignmentInput) (*CreateHospitalToAccountAssignmentResponse, error) {
	assignment, err := s.HospitalSVC.AddHospitalToAccountAssignment(ctx, input.Body.AccountID, input.Body.HospitalID)
	if err != nil {
		return nil, err
	}

	resp := &CreateHospitalToAccountAssignmentResponse{}
	resp.Body.Assignment = *assignment

	return resp, nil
}
