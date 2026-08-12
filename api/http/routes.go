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
		Path:          "/api/v1/accounts/{accountID}",
		Summary:       "/api/v1/accounts/{accountID} - [PUT]",
		Description:   "Updates an existing account",
		Tags:          []string{"accounts"},
		DefaultStatus: http.StatusOK,
	}, s.UpdateAccountHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "get-account",
		Method:        http.MethodGet,
		Path:          "/api/v1/accounts/{accountID}",
		Summary:       "/api/v1/accounts/{accountID} - [GET]",
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

	huma.Register(api, huma.Operation{
		OperationID:   "archive-account",
		Method:        http.MethodDelete,
		Path:          "/api/v1/accounts/{accountID}",
		Summary:       "/api/v1/accounts/{accountID} - [DELETE]",
		Description:   "Archives an existing account",
		Tags:          []string{"accounts"},
		DefaultStatus: http.StatusOK,
	}, s.ArchiveAccountHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "create-doctor",
		Method:        http.MethodPost,
		Path:          "/api/v1/doctors",
		Summary:       "/api/v1/doctors - [POST]",
		Description:   "Creates a new doctor",
		Tags:          []string{"doctors"},
		DefaultStatus: http.StatusCreated,
	}, s.CreateDoctorHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "create-doctor-to-account-assignment",
		Method:        http.MethodPost,
		Path:          "/api/v1/doctors/assignments",
		Summary:       "/api/v1/doctors/assignments - [POST]",
		Description:   "Assigns a doctor to an account",
		Tags:          []string{"doctors"},
		DefaultStatus: http.StatusCreated,
	}, s.CreateDoctorToAccountAssignmentHandler)
}
