package main

import (
	"openhealth/api/http"
	"openhealth/internal/domain"
	"openhealth/internal/pkg/postgres"
	"openhealth/internal/repository"
	"openhealth/internal/service/accounts"
	"os"
)

func main() {
	// POSTGRES
	postgresDB := postgres.NewPostgresDB(os.Getenv("POSTGRES_URI"))
	tx := postgres.NewTxManager(postgresDB.DB)

	// REPOSITORIES
	repository := repository.NewRepository(postgresDB.DB)

	// SERVICES
	accountSVC := accounts.NewService(tx, repository)

	// SERVER
	server := http.NewServer(domain.APIName, domain.APIVersion, os.Getenv("PORT"), accountSVC)
	server.Start()
}
