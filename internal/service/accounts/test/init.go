package test

import (
	"openhealth/internal/service/accounts"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

type Dependencies struct {
	repository *MockRepository
	tx         *MockTxManager
}

// newMockService - creates a new mock service for testing
// Returns the service and its mocked dependencies
func newMockService(t *testing.T) (*accounts.Service, Dependencies) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	tx := NewMockTxManager(ctrl)
	repository := NewMockRepository(ctrl)

	deps := Dependencies{
		repository: repository,
		tx:         tx,
	}

	service := accounts.NewService(tx, repository)

	return service, deps
}
