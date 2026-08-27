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

func TestServer_CreateNurseHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/nurses"

	body := map[string]any{
		"account_id": "mock-account-id-1",
	}

	t.Run("should create nurse", func(t *testing.T) {
		deps.MockNurseSvc.EXPECT().CreateNurse(gomock.Any(), gomock.Any()).Return(nil)

		resp := api.Post(endpoint, addAuthHeader(), body)

		require.Equal(t, http.StatusCreated, resp.Code)

		var response struct {
			Nurse domain.Nurse `json:"nurse"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, "mock-account-id-1", response.Nurse.AccountID)
		require.NotEmpty(t, response.Nurse.ID)
		require.Equal(t, domain.NurseStatusActive, response.Nurse.Status)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockNurseSvc.EXPECT().CreateNurse(gomock.Any(), gomock.Any()).Return(errors.New("mock-error"))

		resp := api.Post(endpoint, addAuthHeader(), body)

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_GetNurseByAccountIDHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/accounts/mock-account-id-1/nurse"

	t.Run("should return nurse for account", func(t *testing.T) {
		nurse := domain.NewNurse("mock-account-id-1")

		deps.MockNurseSvc.EXPECT().GetNurseByAccountID(gomock.Any(), "mock-account-id-1").Return(nurse, nil)

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusOK, resp.Code)

		var response struct {
			Nurse domain.Nurse `json:"nurse"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, *nurse, response.Nurse)
	})

	t.Run("should return 404 when nurse not found", func(t *testing.T) {
		deps.MockNurseSvc.EXPECT().GetNurseByAccountID(gomock.Any(), "mock-account-id-1").Return(nil, domain.NurseNotFoundError)

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockNurseSvc.EXPECT().GetNurseByAccountID(gomock.Any(), "mock-account-id-1").Return(nil, errors.New("mock-error"))

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_GetNurseByIDHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/nurses/mock-nurse-id-1"

	t.Run("should return nurse", func(t *testing.T) {
		nurse := domain.NewNurse("mock-account-id-1")

		deps.MockNurseSvc.EXPECT().GetNurseByID(gomock.Any(), "mock-nurse-id-1").Return(nurse, nil)

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusOK, resp.Code)

		var response struct {
			Nurse domain.Nurse `json:"nurse"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, *nurse, response.Nurse)
	})

	t.Run("should return 404 when nurse not found", func(t *testing.T) {
		deps.MockNurseSvc.EXPECT().GetNurseByID(gomock.Any(), "mock-nurse-id-1").Return(nil, domain.NurseNotFoundError)

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockNurseSvc.EXPECT().GetNurseByID(gomock.Any(), "mock-nurse-id-1").Return(nil, errors.New("mock-error"))

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
