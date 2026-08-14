package authorization

import (
	"context"

	"openhealth/internal/domain"
)

//go:generate mockgen -source=interface.go -destination=test/mocks.go -package=test

type OpenFGA interface {
	AddTuple(ctx context.Context, tuple domain.Tuple) error
	RemoveTuple(ctx context.Context, tuple domain.Tuple) error
}
