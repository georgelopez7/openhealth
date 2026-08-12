package http

import (
	"context"

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
