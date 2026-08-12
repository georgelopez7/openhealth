package http

import "openhealth/internal/domain"

type CreateMedicalRecordInput struct {
	Body struct {
		AccountID   string `json:"account_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}
}

type CreateMedicalRecordResponse struct {
	Body struct {
		MedicalRecord domain.MedicalRecord `json:"medical_record"`
	}
}
