package http

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
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

func (s *Server) GetHospitalByIDHandler(ctx context.Context, input *GetHospitalByIDInput) (*GetHospitalByIDResponse, error) {
	hospital, err := s.HospitalSVC.GetHospitalByID(ctx, input.HospitalID)
	switch err {
	case nil:
		resp := &GetHospitalByIDResponse{}
		resp.Body.Hospital = *hospital
		return resp, nil
	case domain.HospitalNotFoundError:
		return nil, huma.Error404NotFound("hospital not found")
	default:
		return nil, err
	}
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
