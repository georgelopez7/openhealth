package test

import (
	"openhealth/internal/pkg/postgres"
	"openhealth/internal/repository"
	"os"
	"testing"
)

var repo *repository.Repository

func TestMain(m *testing.M) {
	uri, teardown := postgres.SetupMockPostgres()

	db := postgres.NewPostgresDB(uri)

	db.Migrate("../../../_migrations")

	repo = repository.NewRepository(db.DB)

	code := m.Run()

	teardown()

	os.Exit(code)
}
