package test

import (
	"testing"

	xhttp "openhealth/api/http"

	"github.com/danielgtaylor/huma/v2/humatest"
	"go.uber.org/mock/gomock"
)

type Dependencies struct {
	MockAccountSvc *MockAccountSVC
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

	deps := Dependencies{
		MockAccountSvc: mockAccountSvc,
	}

	server := xhttp.NewServer(name, version, port, mockAccountSvc).Mock(t)
	api := server.API.(humatest.TestAPI)

	return api, deps, ctrl.Finish
}
