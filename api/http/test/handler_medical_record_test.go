package test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"openhealth/internal/domain"
)

func TestServer_CreateMedicalRecordHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/medical-records"

	body := map[string]any{
		"account_id":  "mock-account-id-1",
		"title":       "mock-title-1",
		"description": "mock-description-1",
	}

	t.Run("should create medical record", func(t *testing.T) {
		deps.MockMedicalRecordSvc.EXPECT().CreateMedicalRecord(gomock.Any(), gomock.Any()).Return(nil)

		resp := api.Post(endpoint, body)

		require.Equal(t, http.StatusCreated, resp.Code)

		var response struct {
			MedicalRecord domain.MedicalRecord `json:"medical_record"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, "mock-account-id-1", response.MedicalRecord.AccountID)
		require.Equal(t, "mock-title-1", response.MedicalRecord.Title)
		require.Equal(t, "mock-description-1", response.MedicalRecord.Description)
		require.NotEmpty(t, response.MedicalRecord.ID)
		require.Equal(t, domain.MedicalRecordStatusActive, response.MedicalRecord.Status)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockMedicalRecordSvc.EXPECT().CreateMedicalRecord(gomock.Any(), gomock.Any()).Return(errors.New("mock-error"))

		resp := api.Post(endpoint, body)

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
