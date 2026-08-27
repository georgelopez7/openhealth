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

func TestServer_CreateHospitalToAccountAssignmentHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/hospitals/assignments"

	body := map[string]any{
		"account_id":  "mock-account-id-1",
		"hospital_id": "mock-hospital-id-1",
	}

	t.Run("should create hospital to account assignment", func(t *testing.T) {
		deps.MockHospitalSvc.EXPECT().AddHospitalToAccountAssignment(gomock.Any(), "mock-account-id-1", "mock-hospital-id-1").Return(&domain.HospitalAssignment{
			ID:         "mock-assignment-id-1",
			AccountID:  "mock-account-id-1",
			HospitalID: "mock-hospital-id-1",
		}, nil)

		resp := api.Post(endpoint, addAuthHeader(), body)

		require.Equal(t, http.StatusCreated, resp.Code)

		var response struct {
			Assignment domain.HospitalAssignment `json:"assignment"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, "mock-account-id-1", response.Assignment.AccountID)
		require.Equal(t, "mock-hospital-id-1", response.Assignment.HospitalID)
		require.NotEmpty(t, response.Assignment.ID)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockHospitalSvc.EXPECT().AddHospitalToAccountAssignment(gomock.Any(), "mock-account-id-1", "mock-hospital-id-1").Return(nil, errors.New("mock-error"))

		resp := api.Post(endpoint, addAuthHeader(), body)

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
