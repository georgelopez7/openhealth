package http

import (
	"fmt"
	"log/slog"
	"net/http"
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/danielgtaylor/huma/v2/humatest"
)

type Server struct {
	Port             string
	BaseURL          string
	Router           *http.ServeMux
	API              huma.API
	AccountSVC       AccountSVC
	DoctorSVC        DoctorSVC
	NurseSVC         NurseSVC
	HospitalSVC      HospitalSVC
	MedicalRecordSVC MedicalRecordSVC
}

func NewServer(name string, version string, port string, accountSVC AccountSVC, doctorSVC DoctorSVC, nurseSVC NurseSVC, hospitalSVC HospitalSVC, medicalRecordSVC MedicalRecordSVC) *Server {
	router := http.NewServeMux()
	config := huma.DefaultConfig(name, version)

	baseURL := fmt.Sprintf("http://localhost:%s", port)
	config.Servers = []*huma.Server{
		{URL: baseURL},
	}

	api := humago.New(router, config)

	return &Server{
		Port:             port,
		Router:           router,
		API:              api,
		AccountSVC:       accountSVC,
		DoctorSVC:        doctorSVC,
		NurseSVC:         nurseSVC,
		HospitalSVC:      hospitalSVC,
		MedicalRecordSVC: medicalRecordSVC,
	}
}

// AddRoutes - adds the routes to the server
func (s *Server) AddRoutes(api huma.API) {
	s.addRoutes(s.API)
}

// Start - starts the server
func (s *Server) Start() {
	s.addRoutes(s.API)

	port := fmt.Sprintf(":%s", s.Port)

	fmt.Printf("Server listening on port %s\n", port)
	if err := http.ListenAndServe(port, s.Router); err != nil {
		slog.Error("Server failed", "error", err)
	}
}

// Mock - creates a mock instance of the server.
// Replaces the Huma API with a mock Huma API.
func (s *Server) Mock(t *testing.T) *Server {
	_, api := humatest.New(t)

	s.addRoutes(api)

	return &Server{
		API: api,
	}
}
