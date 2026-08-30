package test

import (
	"openhealth/internal/service/outbox"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

type Dependencies struct {
	repository *MockRepository
}

// newMockService - creates a new mock service for testing
// Returns the service and its mocked dependencies
func newMockService(t *testing.T) (*outbox.Service, Dependencies) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	repository := NewMockRepository(ctrl)

	deps := Dependencies{
		repository: repository,
	}

	service := outbox.NewService(repository)

	return service, deps
}
