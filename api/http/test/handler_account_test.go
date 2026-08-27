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

func TestServer_CreateAccountHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/accounts"

	body := map[string]any{
		"first_name": "mock-first-name-1",
		"last_name":  "mock-last-name-1",
		"age":        25,
		"email":      "mock-email-1",
	}

	t.Run("should create account", func(t *testing.T) {
		deps.MockAccountSvc.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(nil)

		resp := api.Post(endpoint, addAuthHeader(), body)

		require.Equal(t, http.StatusCreated, resp.Code)

		var response struct {
			Account domain.Account `json:"account"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, "mock-first-name-1", response.Account.FirstName)
		require.Equal(t, "mock-last-name-1", response.Account.LastName)
		require.Equal(t, 25, response.Account.Age)
		require.Equal(t, "mock-email-1", response.Account.Email)
		require.NotEmpty(t, response.Account.ID)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockAccountSvc.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(errors.New("mock-error"))

		resp := api.Post(endpoint, addAuthHeader(), body)

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_UpdateAccountHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/accounts/mock-id-1"

	body := map[string]any{
		"first_name": "mock-first-name-1",
		"last_name":  "mock-last-name-1",
		"age":        25,
		"email":      "mock-email-1",
	}

	t.Run("should update account", func(t *testing.T) {
		deps.MockAccountSvc.EXPECT().UpdateAccount(gomock.Any(), gomock.Any()).Return(nil)

		resp := api.Put(endpoint, addAuthHeader(), body)

		require.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockAccountSvc.EXPECT().UpdateAccount(gomock.Any(), gomock.Any()).Return(errors.New("mock-error"))

		resp := api.Put(endpoint, addAuthHeader(), body)

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_GetAccountByIDHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/accounts/mock-id-1"

	t.Run("should return account", func(t *testing.T) {
		account := domain.Account{
			ID:        "mock-id-1",
			FirstName: "mock-first-name-1",
			LastName:  "mock-last-name-1",
			Age:       25,
			Email:     "mock-email-1",
		}

		deps.MockAccountSvc.EXPECT().GetAccountByID(gomock.Any(), "mock-id-1").Return(&account, nil)

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusOK, resp.Code)

		var response struct {
			Account domain.Account `json:"account"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, account, response.Account)
	})

	t.Run("should return 404 when account not found", func(t *testing.T) {
		deps.MockAccountSvc.EXPECT().GetAccountByID(gomock.Any(), "mock-id-1").Return(nil, domain.AccountNotFoundError)

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockAccountSvc.EXPECT().GetAccountByID(gomock.Any(), "mock-id-1").Return(nil, errors.New("mock-error"))

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_ListAccountsHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/accounts"

	t.Run("should return accounts", func(t *testing.T) {
		accounts := []domain.Account{
			{
				ID:        "mock-id-1",
				FirstName: "mock-first-name-1",
				LastName:  "mock-last-name-1",
				Age:       25,
				Email:     "mock-email-1",
			},
			{
				ID:        "mock-id-2",
				FirstName: "mock-first-name-2",
				LastName:  "mock-last-name-2",
				Age:       30,
				Email:     "mock-email-2",
			},
		}

		deps.MockAccountSvc.EXPECT().GetAccounts(gomock.Any(), 10).Return(accounts, nil)

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusOK, resp.Code)

		var response struct {
			Accounts []domain.Account `json:"accounts"`
		}

		err := json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)
		require.Equal(t, accounts, response.Accounts)
	})

	t.Run("should pass limit query", func(t *testing.T) {
		deps.MockAccountSvc.EXPECT().GetAccounts(gomock.Any(), 5).Return([]domain.Account{}, nil)

		resp := api.Get(endpoint+"?limit=5", addAuthHeader())

		require.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockAccountSvc.EXPECT().GetAccounts(gomock.Any(), 10).Return(nil, errors.New("mock-error"))

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_ArchiveAccountHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/accounts/mock-id-1"

	t.Run("should archive account", func(t *testing.T) {
		deps.MockAccountSvc.EXPECT().ArchiveAccount(gomock.Any(), "mock-id-1").Return(nil)

		resp := api.Delete(endpoint, addAuthHeader())

		require.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("should handle service error", func(t *testing.T) {
		deps.MockAccountSvc.EXPECT().ArchiveAccount(gomock.Any(), "mock-id-1").Return(errors.New("mock-error"))

		resp := api.Delete(endpoint, addAuthHeader())

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
