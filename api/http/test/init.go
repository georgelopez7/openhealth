package test

import (
	"testing"

	xhttp "openhealth/api/http"

	"github.com/danielgtaylor/huma/v2/humatest"
	"go.uber.org/mock/gomock"
)

type Dependencies struct {
	MockAccountSvc *MockAccountSVC
	MockDoctorSvc  *MockDoctorSVC
}

// newMockServer - creates a mock server.
func newMockServer(t *testing.T) (humatest.TestAPI, Dependencies, func()) {
	ctrl := gomock.NewController(t)

	const (
		name    = "test-server"
		version = "1.0.0"
		port    = "8080"
	)

	mockAccountSvc := NewMockAccountSVC(ctrl)
	mockDoctorSvc := NewMockDoctorSVC(ctrl)

	deps := Dependencies{
		MockAccountSvc: mockAccountSvc,
		MockDoctorSvc:  mockDoctorSvc,
	}

	server := xhttp.NewServer(name, version, port, mockAccountSvc, mockDoctorSvc).Mock(t)
	api := server.API.(humatest.TestAPI)

	return api, deps, ctrl.Finish
}
