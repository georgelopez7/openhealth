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

func TestServer_CreateHospitalHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/hospitals"

	body := map[string]any{
		"name": "mock-hospital-name-1",
	}

	t.Run("should create hospital", func(t *testing.T) {
		deps.MockHospitalSvc.EXPECT().CreateHospital(gomock.Any(), gomock.Any()).Return(nil)

		resp := api.Post(endpoint, body)

		require.Equal(t, http.StatusCreated, resp.Code)

		var response struct {
			Hospital domain.Hospital `json:"hospital"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, "mock-hospital-name-1", response.Hospital.Name)
		require.NotEmpty(t, response.Hospital.ID)
		require.Equal(t, domain.HospitalStatusActive, response.Hospital.Status)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockHospitalSvc.EXPECT().CreateHospital(gomock.Any(), gomock.Any()).Return(errors.New("mock-error"))

		resp := api.Post(endpoint, body)

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
