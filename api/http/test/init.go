package test

import (
	"testing"

	xhttp "openhealth/api/http"

	"github.com/danielgtaylor/huma/v2/humatest"
	"go.uber.org/mock/gomock"
)

type Dependencies struct {
	MockAccountSvc       *MockAccountSVC
	MockDoctorSvc        *MockDoctorSVC
	MockNurseSvc         *MockNurseSVC
	MockHospitalSvc      *MockHospitalSVC
	MockMedicalRecordSvc *MockMedicalRecordSVC
}

// newMockServer - creates a mock server.
func newMockServer(t *testing.T) (humatest.TestAPI, Dependencies, func()) {
	ctrl := gomock.NewController(t)

	const (
		name      = "test-server"
		version   = "1.0.0"
		port      = "8080"
		authToken = "test-token"
	)

	mockAccountSvc := NewMockAccountSVC(ctrl)
	mockDoctorSvc := NewMockDoctorSVC(ctrl)
	mockNurseSvc := NewMockNurseSVC(ctrl)
	mockHospitalSvc := NewMockHospitalSVC(ctrl)
	mockMedicalRecordSvc := NewMockMedicalRecordSVC(ctrl)

	deps := Dependencies{
		MockAccountSvc:       mockAccountSvc,
		MockDoctorSvc:        mockDoctorSvc,
		MockNurseSvc:         mockNurseSvc,
		MockHospitalSvc:      mockHospitalSvc,
		MockMedicalRecordSvc: mockMedicalRecordSvc,
	}

	server := xhttp.NewServer(name, version, port, authToken, mockAccountSvc, mockDoctorSvc, mockNurseSvc, mockHospitalSvc, mockMedicalRecordSvc).Mock(t)
	api := server.API.(humatest.TestAPI)

	return api, deps, ctrl.Finish
}

// addAuthHeader - returns the Authorization header value used by the test server.
func addAuthHeader() string {
	return "Authorization: Bearer test-token"
}
