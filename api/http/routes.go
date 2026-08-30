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
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.CreateAccountHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "update-account",
		Method:        http.MethodPut,
		Path:          "/api/v1/accounts/{accountID}",
		Summary:       "/api/v1/accounts/{accountID} - [PUT]",
		Description:   "Updates an existing account",
		Tags:          []string{"accounts"},
		DefaultStatus: http.StatusOK,
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.UpdateAccountHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "get-account",
		Method:        http.MethodGet,
		Path:          "/api/v1/accounts/{accountID}",
		Summary:       "/api/v1/accounts/{accountID} - [GET]",
		Description:   "Gets an account by its ID",
		Tags:          []string{"accounts"},
		DefaultStatus: http.StatusOK,
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.GetAccountByIDHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "list-accounts",
		Method:        http.MethodGet,
		Path:          "/api/v1/accounts",
		Summary:       "/api/v1/accounts - [GET]",
		Description:   "Lists accounts up to the provided limit",
		Tags:          []string{"accounts"},
		DefaultStatus: http.StatusOK,
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.ListAccountsHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "archive-account",
		Method:        http.MethodDelete,
		Path:          "/api/v1/accounts/{accountID}",
		Summary:       "/api/v1/accounts/{accountID} - [DELETE]",
		Description:   "Archives an existing account",
		Tags:          []string{"accounts"},
		DefaultStatus: http.StatusOK,
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.ArchiveAccountHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "create-doctor",
		Method:        http.MethodPost,
		Path:          "/api/v1/doctors",
		Summary:       "/api/v1/doctors - [POST]",
		Description:   "Creates a new doctor",
		Tags:          []string{"doctors"},
		DefaultStatus: http.StatusCreated,
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.CreateDoctorHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "get-doctor",
		Method:        http.MethodGet,
		Path:          "/api/v1/doctors/{doctorID}",
		Summary:       "/api/v1/doctors/{doctorID} - [GET]",
		Description:   "Gets a doctor by its ID",
		Tags:          []string{"doctors"},
		DefaultStatus: http.StatusOK,
		Errors:        []int{http.StatusNotFound},
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.GetDoctorByIDHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "get-doctor-by-account",
		Method:        http.MethodGet,
		Path:          "/api/v1/accounts/{accountID}/doctor",
		Summary:       "/api/v1/accounts/{accountID}/doctor - [GET]",
		Description:   "Gets the doctor associated with an account",
		Tags:          []string{"doctors"},
		DefaultStatus: http.StatusOK,
		Errors:        []int{http.StatusNotFound},
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.GetDoctorByAccountIDHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "create-doctor-to-account-assignment",
		Method:        http.MethodPost,
		Path:          "/api/v1/doctors/assignments",
		Summary:       "/api/v1/doctors/assignments - [POST]",
		Description:   "Assigns a doctor to an account",
		Tags:          []string{"doctors"},
		DefaultStatus: http.StatusCreated,
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.CreateDoctorToAccountAssignmentHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "create-nurse",
		Method:        http.MethodPost,
		Path:          "/api/v1/nurses",
		Summary:       "/api/v1/nurses - [POST]",
		Description:   "Creates a new nurse",
		Tags:          []string{"nurses"},
		DefaultStatus: http.StatusCreated,
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.CreateNurseHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "get-nurse",
		Method:        http.MethodGet,
		Path:          "/api/v1/nurses/{nurseID}",
		Summary:       "/api/v1/nurses/{nurseID} - [GET]",
		Description:   "Gets a nurse by its ID",
		Tags:          []string{"nurses"},
		DefaultStatus: http.StatusOK,
		Errors:        []int{http.StatusNotFound},
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.GetNurseByIDHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "get-nurse-by-account",
		Method:        http.MethodGet,
		Path:          "/api/v1/accounts/{accountID}/nurse",
		Summary:       "/api/v1/accounts/{accountID}/nurse - [GET]",
		Description:   "Gets the nurse associated with an account",
		Tags:          []string{"nurses"},
		DefaultStatus: http.StatusOK,
		Errors:        []int{http.StatusNotFound},
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.GetNurseByAccountIDHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "create-nurse-to-hospital-assignment",
		Method:        http.MethodPost,
		Path:          "/api/v1/nurses/assignments",
		Summary:       "/api/v1/nurses/assignments - [POST]",
		Description:   "Assigns a nurse to a hospital",
		Tags:          []string{"nurses"},
		DefaultStatus: http.StatusCreated,
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.CreateNurseToHospitalAssignmentHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "create-hospital",
		Method:        http.MethodPost,
		Path:          "/api/v1/hospitals",
		Summary:       "/api/v1/hospitals - [POST]",
		Description:   "Creates a new hospital",
		Tags:          []string{"hospitals"},
		DefaultStatus: http.StatusCreated,
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.CreateHospitalHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "get-hospital",
		Method:        http.MethodGet,
		Path:          "/api/v1/hospitals/{hospitalID}",
		Summary:       "/api/v1/hospitals/{hospitalID} - [GET]",
		Description:   "Gets a hospital by its ID",
		Tags:          []string{"hospitals"},
		DefaultStatus: http.StatusOK,
		Errors:        []int{http.StatusNotFound},
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.GetHospitalByIDHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "create-hospital-to-account-assignment",
		Method:        http.MethodPost,
		Path:          "/api/v1/hospitals/assignments",
		Summary:       "/api/v1/hospitals/assignments - [POST]",
		Description:   "Assigns a hospital to an account",
		Tags:          []string{"hospitals"},
		DefaultStatus: http.StatusCreated,
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.CreateHospitalToAccountAssignmentHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "create-medical-record",
		Method:        http.MethodPost,
		Path:          "/api/v1/medical-records",
		Summary:       "/api/v1/medical-records - [POST]",
		Description:   "Creates a new medical record",
		Tags:          []string{"medical-records"},
		DefaultStatus: http.StatusCreated,
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.CreateMedicalRecordHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "get-medical-record",
		Method:        http.MethodGet,
		Path:          "/api/v1/medical-records/{medicalRecordID}",
		Summary:       "/api/v1/medical-records/{medicalRecordID} - [GET]",
		Description:   "Gets a medical record by its ID",
		Tags:          []string{"medical-records"},
		DefaultStatus: http.StatusOK,
		Errors:        []int{http.StatusNotFound},
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.GetMedicalRecordByIDHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "get-medical-record-access",
		Method:        http.MethodGet,
		Path:          "/api/v1/accounts/{accountID}/medical-records/{medicalRecordID}/access",
		Summary:       "/api/v1/accounts/{accountID}/medical-records/{medicalRecordID}/access - [GET]",
		Description:   "Checks whether an account can view and edit a medical record",
		Tags:          []string{"medical-records"},
		DefaultStatus: http.StatusOK,
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.GetMedicalRecordAccessHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "list-medical-records",
		Method:        http.MethodGet,
		Path:          "/api/v1/medical-records",
		Summary:       "/api/v1/medical-records - [GET]",
		Description:   "Lists all medical records",
		Tags:          []string{"medical-records"},
		DefaultStatus: http.StatusOK,
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.GetMedicalRecordsHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "list-medical-records-by-account",
		Method:        http.MethodGet,
		Path:          "/api/v1/accounts/{accountID}/medical-records",
		Summary:       "/api/v1/accounts/{accountID}/medical-records - [GET]",
		Description:   "Lists medical records for an account",
		Tags:          []string{"medical-records"},
		DefaultStatus: http.StatusOK,
		Middlewares:   huma.Middlewares{s.ValidateAuthTokenMiddleware},
	}, s.GetMedicalRecordsByAccountIDHandler)
}
