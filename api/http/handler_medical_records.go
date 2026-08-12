package http

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"openhealth/internal/domain"
)

func (s *Server) CreateMedicalRecordHandler(ctx context.Context, input *CreateMedicalRecordInput) (*CreateMedicalRecordResponse, error) {
	record := domain.NewMedicalRecord(
		input.Body.AccountID,
		input.Body.Title,
		input.Body.Description,
	)

	err := s.MedicalRecordSVC.CreateMedicalRecord(ctx, *record)
	if err != nil {
		return nil, err
	}

	resp := &CreateMedicalRecordResponse{}
	resp.Body.MedicalRecord = *record

	return resp, nil
}

func (s *Server) GetMedicalRecordByIDHandler(ctx context.Context, input *GetMedicalRecordByIDInput) (*GetMedicalRecordByIDResponse, error) {
	record, err := s.MedicalRecordSVC.GetMedicalRecordByID(ctx, input.MedicalRecordID)
	switch err {
	case nil:
		resp := &GetMedicalRecordByIDResponse{}
		resp.Body.MedicalRecord = *record
		return resp, nil
	case domain.MedicalRecordNotFoundError:
		return nil, huma.Error404NotFound("medical record not found")
	default:
		return nil, err
	}
}

func (s *Server) GetMedicalRecordsByAccountIDHandler(ctx context.Context, input *GetMedicalRecordsByAccountIDInput) (*GetMedicalRecordsByAccountIDResponse, error) {
	records, err := s.MedicalRecordSVC.GetMedicalRecordsByAccountID(ctx, input.AccountID)
	if err != nil {
		return nil, err
	}

	resp := &GetMedicalRecordsByAccountIDResponse{}
	resp.Body.MedicalRecords = records

	return resp, nil
}
