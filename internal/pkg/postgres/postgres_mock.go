package postgres

import (
	"context"
	"log"

	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

// SetupMockPostgres - starts a mock postgres container
func SetupMockPostgres() (string, func()) {
	ctx := context.Background()

	postgresContainer, err := postgres.Run(ctx,
		"postgres:18-alpine",
		postgres.WithDatabase("test"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		log.Fatalf("failed to start mock postgres container: %v", err)
	}

	teardown := func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate mock postgres container: %v", err)
		}
	}

	postgresURI, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalf("failed to get mock postgres connection string: %v", err)
	}

	return postgresURI, teardown
}
