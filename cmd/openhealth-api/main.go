package main

import (
	"openhealth/api/http"
	"openhealth/internal/domain"
	"openhealth/internal/pkg/postgres"
	"openhealth/internal/repository"
	"openhealth/internal/service/accounts"
	"openhealth/internal/service/doctors"
	"openhealth/internal/service/hospitals"
	"openhealth/internal/service/nurses"
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
	doctorSVC := doctors.NewService(tx, repository)
	nurseSVC := nurses.NewService(tx, repository)
	hospitalSVC := hospitals.NewService(tx, repository)

	// SERVER
	server := http.NewServer(domain.APIName, domain.APIVersion, os.Getenv("PORT"), accountSVC, doctorSVC, nurseSVC, hospitalSVC)
	server.Start()
}
