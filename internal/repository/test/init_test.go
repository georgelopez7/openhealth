package test

import (
	"openhealth/internal/pkg/postgres"
	"openhealth/internal/repository"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
)

var (
	repo *repository.Repository
	db   *sqlx.DB
)

func TestMain(m *testing.M) {
	uri, teardown := postgres.SetupMockPostgres()

	postgresDB := postgres.NewPostgresDB(uri)

	postgresDB.Migrate("../../../_migrations")

	db = postgresDB.DB

	repo = repository.NewRepository(db)

	code := m.Run()

	teardown()

	os.Exit(code)
}
