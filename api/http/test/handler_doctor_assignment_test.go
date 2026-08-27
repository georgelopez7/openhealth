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

func TestServer_CreateDoctorToAccountAssignmentHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/doctors/assignments"

	body := map[string]any{
		"account_id": "mock-account-id-1",
		"doctor_id":  "mock-doctor-id-1",
	}

	t.Run("should create doctor to account assignment", func(t *testing.T) {
		deps.MockDoctorSvc.EXPECT().AddDoctorToAccountAssignment(gomock.Any(), "mock-account-id-1", "mock-doctor-id-1").Return(&domain.DoctorAssignment{
			ID:        "mock-assignment-id-1",
			AccountID: "mock-account-id-1",
			DoctorID:  "mock-doctor-id-1",
		}, nil)

		resp := api.Post(endpoint, addAuthHeader(), body)

		require.Equal(t, http.StatusCreated, resp.Code)

		var response struct {
			Assignment domain.DoctorAssignment `json:"assignment"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, "mock-account-id-1", response.Assignment.AccountID)
		require.Equal(t, "mock-doctor-id-1", response.Assignment.DoctorID)
		require.NotEmpty(t, response.Assignment.ID)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockDoctorSvc.EXPECT().AddDoctorToAccountAssignment(gomock.Any(), "mock-account-id-1", "mock-doctor-id-1").Return(nil, errors.New("mock-error"))

		resp := api.Post(endpoint, addAuthHeader(), body)

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
