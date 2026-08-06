package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func (s *Server) addRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "healthcheck",
		Method:        http.MethodGet,
		Path:          "/api/v1/healthcheck",
		Summary:       "/api/v1/healthcheck - [GET]",
		Description:   "Checks if the service is healthy",
		Tags:          []string{"system"},
		DefaultStatus: http.StatusOK,
	}, s.HealthcheckHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "create-account",
		Method:        http.MethodPost,
		Path:          "/api/v1/accounts",
		Summary:       "/api/v1/accounts - [POST]",
		Description:   "Creates a new account",
		Tags:          []string{"accounts"},
		DefaultStatus: http.StatusCreated,
	}, s.CreateAccountHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "update-account",
		Method:        http.MethodPut,
		Path:          "/api/v1/accounts/{id}",
		Summary:       "/api/v1/accounts/{id} - [PUT]",
		Description:   "Updates an existing account",
		Tags:          []string{"accounts"},
		DefaultStatus: http.StatusOK,
	}, s.UpdateAccountHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "get-account",
		Method:        http.MethodGet,
		Path:          "/api/v1/accounts/{id}",
		Summary:       "/api/v1/accounts/{id} - [GET]",
		Description:   "Gets an account by its ID",
		Tags:          []string{"accounts"},
		DefaultStatus: http.StatusOK,
	}, s.GetAccountByIDHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "list-accounts",
		Method:        http.MethodGet,
		Path:          "/api/v1/accounts",
		Summary:       "/api/v1/accounts - [GET]",
		Description:   "Lists accounts up to the provided limit",
		Tags:          []string{"accounts"},
		DefaultStatus: http.StatusOK,
	}, s.ListAccountsHandler)
}
