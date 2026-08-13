package http

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
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

func (s *Server) GetDoctorByIDHandler(ctx context.Context, input *GetDoctorByIDInput) (*GetDoctorByIDResponse, error) {
	doctor, err := s.DoctorSVC.GetDoctorByID(ctx, input.DoctorID)
	switch err {
	case nil:
		resp := &GetDoctorByIDResponse{}
		resp.Body.Doctor = *doctor
		return resp, nil
	case domain.DoctorNotFoundError:
		return nil, huma.Error404NotFound("doctor not found")
	default:
		return nil, err
	}
}

func (s *Server) GetDoctorByAccountIDHandler(ctx context.Context, input *GetDoctorByAccountIDInput) (*GetDoctorByAccountIDResponse, error) {
	doctor, err := s.DoctorSVC.GetDoctorByAccountID(ctx, input.AccountID)
	switch err {
	case nil:
		resp := &GetDoctorByAccountIDResponse{}
		resp.Body.Doctor = *doctor
		return resp, nil
	case domain.DoctorNotFoundError:
		return nil, huma.Error404NotFound("doctor not found")
	default:
		return nil, err
	}
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
