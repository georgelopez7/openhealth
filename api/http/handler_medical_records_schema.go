package http

import "openhealth/internal/domain"

type CreateMedicalRecordInput struct {
	Body struct {
		AccountID   string `json:"account_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}
}

type GetMedicalRecordByIDInput struct {
	MedicalRecordID string `path:"medicalRecordID"`
}

type GetMedicalRecordByIDResponse struct {
	Body struct {
		MedicalRecord domain.MedicalRecord `json:"medical_record"`
	}
}

type GetMedicalRecordsInput struct{}

type GetMedicalRecordsResponse struct {
	Body struct {
		MedicalRecords []domain.MedicalRecord `json:"medical_records"`
	}
}

type GetMedicalRecordsByAccountIDInput struct {
	AccountID string `path:"accountID"`
}

type GetMedicalRecordsByAccountIDResponse struct {
	Body struct {
		MedicalRecords []domain.MedicalRecord `json:"medical_records"`
	}
}

type CreateMedicalRecordResponse struct {
	Body struct {
		MedicalRecord domain.MedicalRecord `json:"medical_record"`
	}
}
