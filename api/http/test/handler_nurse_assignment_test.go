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

func TestServer_CreateNurseToHospitalAssignmentHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/nurses/assignments"

	body := map[string]any{
		"nurse_id":    "mock-nurse-id-1",
		"hospital_id": "mock-hospital-id-1",
	}

	t.Run("should create nurse to hospital assignment", func(t *testing.T) {
		deps.MockNurseSvc.EXPECT().AddNurseToHospitalAssignment(gomock.Any(), "mock-nurse-id-1", "mock-hospital-id-1").Return(&domain.NurseToHospitalAssignment{
			ID:         "mock-assignment-id-1",
			NurseID:    "mock-nurse-id-1",
			HospitalID: "mock-hospital-id-1",
		}, nil)

		resp := api.Post(endpoint, addAuthHeader(), body)

		require.Equal(t, http.StatusCreated, resp.Code)

		var response struct {
			Assignment domain.NurseToHospitalAssignment `json:"assignment"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, "mock-nurse-id-1", response.Assignment.NurseID)
		require.Equal(t, "mock-hospital-id-1", response.Assignment.HospitalID)
		require.NotEmpty(t, response.Assignment.ID)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockNurseSvc.EXPECT().AddNurseToHospitalAssignment(gomock.Any(), "mock-nurse-id-1", "mock-hospital-id-1").Return(nil, errors.New("mock-error"))

		resp := api.Post(endpoint, addAuthHeader(), body)

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
