package http

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"openhealth/internal/domain"
)

func (s *Server) CreateNurseHandler(ctx context.Context, input *CreateNurseInput) (*CreateNurseResponse, error) {
	nurse := domain.NewNurse(input.Body.AccountID)

	err := s.NurseSVC.CreateNurse(ctx, *nurse)
	if err != nil {
		return nil, err
	}

	resp := &CreateNurseResponse{}
	resp.Body.Nurse = *nurse

	return resp, nil
}

func (s *Server) GetNurseByIDHandler(ctx context.Context, input *GetNurseByIDInput) (*GetNurseByIDResponse, error) {
	nurse, err := s.NurseSVC.GetNurseByID(ctx, input.NurseID)
	switch err {
	case nil:
		resp := &GetNurseByIDResponse{}
		resp.Body.Nurse = *nurse
		return resp, nil
	case domain.NurseNotFoundError:
		return nil, huma.Error404NotFound("nurse not found")
	default:
		return nil, err
	}
}

func (s *Server) GetNurseByAccountIDHandler(ctx context.Context, input *GetNurseByAccountIDInput) (*GetNurseByAccountIDResponse, error) {
	nurse, err := s.NurseSVC.GetNurseByAccountID(ctx, input.AccountID)
	switch err {
	case nil:
		resp := &GetNurseByAccountIDResponse{}
		resp.Body.Nurse = *nurse
		return resp, nil
	case domain.NurseNotFoundError:
		return nil, huma.Error404NotFound("nurse not found")
	default:
		return nil, err
	}
}

func (s *Server) CreateNurseToHospitalAssignmentHandler(ctx context.Context, input *CreateNurseToHospitalAssignmentInput) (*CreateNurseToHospitalAssignmentResponse, error) {
	assignment, err := s.NurseSVC.AddNurseToHospitalAssignment(ctx, input.Body.NurseID, input.Body.HospitalID)
	if err != nil {
		return nil, err
	}

	resp := &CreateNurseToHospitalAssignmentResponse{}
	resp.Body.Assignment = *assignment

	return resp, nil
}
