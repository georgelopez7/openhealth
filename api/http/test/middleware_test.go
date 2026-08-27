package test

import (
	"net/http"
	"testing"

	"openhealth/internal/domain"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestMiddleware_ValidateAuthTokenMiddleware(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	endpoint := "/api/v1/accounts"

	t.Run("should return 401 when the authorization header is missing", func(t *testing.T) {
		resp := api.Get(endpoint)

		require.Equal(t, http.StatusUnauthorized, resp.Code)
	})

	t.Run("should return 401 when the authorization header is not a bearer token", func(t *testing.T) {
		resp := api.Get(endpoint, "Authorization: Basic dXNlcjpwYXNz")

		require.Equal(t, http.StatusUnauthorized, resp.Code)
	})

	t.Run("should return 401 when the token is invalid", func(t *testing.T) {
		resp := api.Get(endpoint, "Authorization: Bearer wrong-token")

		require.Equal(t, http.StatusUnauthorized, resp.Code)
	})

	t.Run("should allow the request when the token is valid", func(t *testing.T) {
		deps.MockAccountSvc.EXPECT().GetAccounts(gomock.Any(), gomock.Any()).Return([]domain.Account{}, nil)

		resp := api.Get(endpoint, addAuthHeader())

		require.Equal(t, http.StatusOK, resp.Code)
	})
}
