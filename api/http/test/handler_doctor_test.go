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

func TestServer_CreateDoctorHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/doctors"

	body := map[string]any{
		"account_id": "mock-account-id-1",
	}

	t.Run("should create doctor", func(t *testing.T) {
		deps.MockDoctorSvc.EXPECT().CreateDoctor(gomock.Any(), gomock.Any()).Return(nil)

		resp := api.Post(endpoint, body)

		require.Equal(t, http.StatusCreated, resp.Code)

		var response struct {
			Doctor domain.Doctor `json:"doctor"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, "mock-account-id-1", response.Doctor.AccountID)
		require.NotEmpty(t, response.Doctor.ID)
		require.Equal(t, domain.DoctorStatusActive, response.Doctor.Status)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockDoctorSvc.EXPECT().CreateDoctor(gomock.Any(), gomock.Any()).Return(errors.New("mock-error"))

		resp := api.Post(endpoint, body)

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_GetDoctorByAccountIDHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/accounts/mock-account-id-1/doctor"

	t.Run("should return doctor for account", func(t *testing.T) {
		doctor := domain.NewDoctor("mock-account-id-1")

		deps.MockDoctorSvc.EXPECT().GetDoctorByAccountID(gomock.Any(), "mock-account-id-1").Return(doctor, nil)

		resp := api.Get(endpoint)

		require.Equal(t, http.StatusOK, resp.Code)

		var response struct {
			Doctor domain.Doctor `json:"doctor"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, *doctor, response.Doctor)
	})

	t.Run("should return 404 when doctor not found", func(t *testing.T) {
		deps.MockDoctorSvc.EXPECT().GetDoctorByAccountID(gomock.Any(), "mock-account-id-1").Return(nil, domain.DoctorNotFoundError)

		resp := api.Get(endpoint)

		require.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockDoctorSvc.EXPECT().GetDoctorByAccountID(gomock.Any(), "mock-account-id-1").Return(nil, errors.New("mock-error"))

		resp := api.Get(endpoint)

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_GetDoctorByIDHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/doctors/mock-doctor-id-1"

	t.Run("should return doctor", func(t *testing.T) {
		doctor := domain.NewDoctor("mock-account-id-1")

		deps.MockDoctorSvc.EXPECT().GetDoctorByID(gomock.Any(), "mock-doctor-id-1").Return(doctor, nil)

		resp := api.Get(endpoint)

		require.Equal(t, http.StatusOK, resp.Code)

		var response struct {
			Doctor domain.Doctor `json:"doctor"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, *doctor, response.Doctor)
	})

	t.Run("should return 404 when doctor not found", func(t *testing.T) {
		deps.MockDoctorSvc.EXPECT().GetDoctorByID(gomock.Any(), "mock-doctor-id-1").Return(nil, domain.DoctorNotFoundError)

		resp := api.Get(endpoint)

		require.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockDoctorSvc.EXPECT().GetDoctorByID(gomock.Any(), "mock-doctor-id-1").Return(nil, errors.New("mock-error"))

		resp := api.Get(endpoint)

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
