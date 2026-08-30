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

		resp := api.Post(endpoint, addAuthHeader(), body)

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

		resp := api.Post(endpoint, addAuthHeader(), body)

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_GetMedicalRecordByIDHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/medical-records/mock-record-id-1"

	t.Run("should return medical record", func(t *testing.T) {
		record := domain.NewMedicalRecord("mock-account-id-1", "mock-title-1", "mock-description-1")

		deps.MockMedicalRecordSvc.EXPECT().GetMedicalRecordByID(gomock.Any(), "mock-record-id-1").Return(record, nil)

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusOK, resp.Code)

		var response struct {
			MedicalRecord domain.MedicalRecord `json:"medical_record"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, *record, response.MedicalRecord)
	})

	t.Run("should return 404 when medical record not found", func(t *testing.T) {
		deps.MockMedicalRecordSvc.EXPECT().GetMedicalRecordByID(gomock.Any(), "mock-record-id-1").Return(nil, domain.MedicalRecordNotFoundError)

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockMedicalRecordSvc.EXPECT().GetMedicalRecordByID(gomock.Any(), "mock-record-id-1").Return(nil, errors.New("mock-error"))

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_GetMedicalRecordsHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/medical-records"

	t.Run("should return all medical records", func(t *testing.T) {
		records := []domain.MedicalRecord{
			*domain.NewMedicalRecord("mock-account-id-1", "Record 1", "First record"),
			*domain.NewMedicalRecord("mock-account-id-2", "Record 2", "Second record"),
		}

		deps.MockMedicalRecordSvc.EXPECT().GetMedicalRecords(gomock.Any()).Return(records, nil)

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusOK, resp.Code)

		var response struct {
			MedicalRecords []domain.MedicalRecord `json:"medical_records"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, records, response.MedicalRecords)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockMedicalRecordSvc.EXPECT().GetMedicalRecords(gomock.Any()).Return(nil, errors.New("mock-error"))

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_GetMedicalRecordsByAccountIDHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/accounts/mock-account-id-1/medical-records"

	t.Run("should return medical records for account", func(t *testing.T) {
		records := []domain.MedicalRecord{
			*domain.NewMedicalRecord("mock-account-id-1", "Record 1", "First record"),
			*domain.NewMedicalRecord("mock-account-id-1", "Record 2", "Second record"),
		}

		deps.MockMedicalRecordSvc.EXPECT().GetMedicalRecordsByAccountID(gomock.Any(), "mock-account-id-1").Return(records, nil)

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusOK, resp.Code)

		var response struct {
			MedicalRecords []domain.MedicalRecord `json:"medical_records"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, records, response.MedicalRecords)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockMedicalRecordSvc.EXPECT().GetMedicalRecordsByAccountID(gomock.Any(), "mock-account-id-1").Return(nil, errors.New("mock-error"))

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_GetMedicalRecordAccessHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/accounts/mock-account-id-1/medical-records/mock-record-id-1/access"

	t.Run("should return access flags", func(t *testing.T) {
		deps.MockMedicalRecordSvc.EXPECT().GetMedicalRecordAccess(gomock.Any(), "mock-record-id-1", "mock-account-id-1").Return(true, false, nil)

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusOK, resp.Code)

		var response struct {
			CanView bool `json:"can_view"`
			CanEdit bool `json:"can_edit"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.True(t, response.CanView)
		require.False(t, response.CanEdit)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockMedicalRecordSvc.EXPECT().GetMedicalRecordAccess(gomock.Any(), "mock-record-id-1", "mock-account-id-1").Return(false, false, errors.New("mock-error"))

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
