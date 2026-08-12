package test

import (
	"testing"

	xhttp "openhealth/api/http"

	"github.com/danielgtaylor/huma/v2/humatest"
	"go.uber.org/mock/gomock"
)

type Dependencies struct {
	MockAccountSvc  *MockAccountSVC
	MockDoctorSvc   *MockDoctorSVC
	MockNurseSvc    *MockNurseSVC
	MockHospitalSvc *MockHospitalSVC
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
	mockNurseSvc := NewMockNurseSVC(ctrl)
	mockHospitalSvc := NewMockHospitalSVC(ctrl)

	deps := Dependencies{
		MockAccountSvc:  mockAccountSvc,
		MockDoctorSvc:   mockDoctorSvc,
		MockNurseSvc:    mockNurseSvc,
		MockHospitalSvc: mockHospitalSvc,
	}

	server := xhttp.NewServer(name, version, port, mockAccountSvc, mockDoctorSvc, mockNurseSvc, mockHospitalSvc).Mock(t)
	api := server.API.(humatest.TestAPI)

	return api, deps, ctrl.Finish
}
