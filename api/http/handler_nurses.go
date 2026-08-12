package http

import (
	"context"

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

func (s *Server) CreateNurseToHospitalAssignmentHandler(ctx context.Context, input *CreateNurseToHospitalAssignmentInput) (*CreateNurseToHospitalAssignmentResponse, error) {
	assignment, err := s.NurseSVC.AddNurseToHospitalAssignment(ctx, input.Body.NurseID, input.Body.HospitalID)
	if err != nil {
		return nil, err
	}

	resp := &CreateNurseToHospitalAssignmentResponse{}
	resp.Body.Assignment = *assignment

	return resp, nil
}
