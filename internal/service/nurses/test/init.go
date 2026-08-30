package test

import (
	"openhealth/internal/service/nurses"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

type Dependencies struct {
	repository *MockRepository
	tx         *MockTxManager
}

// newMockService - creates a new mock service for testing
func newMockService(t *testing.T) (*nurses.Service, Dependencies) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	tx := NewMockTxManager(ctrl)
	repository := NewMockRepository(ctrl)

	deps := Dependencies{
		repository: repository,
		tx:         tx,
	}

	service := nurses.NewService(tx, repository)

	return service, deps
}
