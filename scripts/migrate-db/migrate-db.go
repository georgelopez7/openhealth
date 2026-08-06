package main

import (
	"os"

	"openhealth/internal/pkg/postgres"
)

func main() {
	postgresDB := postgres.NewPostgresDB(os.Getenv("POSTGRES_URI"))
	postgresDB.Migrate("_migrations")
}
