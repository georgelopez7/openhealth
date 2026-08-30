package test

import (
	"openhealth/internal/service/medical-records"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

type Dependencies struct {
	repository       *MockRepository
	tx               *MockTxManager
	authorizationSVC *MockAuthorizationSVC
}

// newMockService - creates a new mock service for testing
func newMockService(t *testing.T) (*medicalrecords.Service, Dependencies) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	tx := NewMockTxManager(ctrl)
	repository := NewMockRepository(ctrl)
	authorizationSVC := NewMockAuthorizationSVC(ctrl)

	deps := Dependencies{
		repository:       repository,
		tx:               tx,
		authorizationSVC: authorizationSVC,
	}

	service := medicalrecords.NewService(tx, repository, authorizationSVC)

	return service, deps
}
