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
