package http

import (
	"context"

	"openhealth/internal/domain"
)

func (s *Server) CreateDoctorHandler(ctx context.Context, input *CreateDoctorInput) (*CreateDoctorResponse, error) {
	doctor := domain.NewDoctor(input.Body.AccountID)

	err := s.DoctorSVC.CreateDoctor(ctx, *doctor)
	if err != nil {
		return nil, err
	}

	resp := &CreateDoctorResponse{}
	resp.Body.Doctor = *doctor

	return resp, nil
}

func (s *Server) CreateDoctorToAccountAssignmentHandler(ctx context.Context, input *CreateDoctorToAccountAssignmentInput) (*CreateDoctorToAccountAssignmentResponse, error) {
	assignment, err := s.DoctorSVC.AddDoctorToAccountAssignment(ctx, input.Body.AccountID, input.Body.DoctorID)
	if err != nil {
		return nil, err
	}

	resp := &CreateDoctorToAccountAssignmentResponse{}
	resp.Body.Assignment = *assignment

	return resp, nil
}
