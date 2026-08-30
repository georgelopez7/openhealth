package main

import (
	"context"
	"log"
	"openhealth/api/http"
	"openhealth/internal/domain"
	"openhealth/internal/pkg/postgres"
	"openhealth/internal/relay"
	"openhealth/internal/repository"
	"openhealth/internal/service/accounts"
	"openhealth/internal/service/authorization"
	"openhealth/internal/service/doctors"
	"openhealth/internal/service/hospitals"
	medicalrecords "openhealth/internal/service/medical-records"
	"openhealth/internal/service/nurses"
	"openhealth/internal/service/outbox"
	"openhealth/pkg/openfga"
	"os"
	"time"
)

func main() {
	// CONFIGURATION
	ctx := context.Background()

	if os.Getenv("FGA_STORE_ID") == "" || os.Getenv("FGA_MODEL_ID") == "" {
		log.Fatal("Please set FGA_STORE_ID and FGA_MODEL_ID environment variables.")
	}

	// POSTGRES
	postgresDB := postgres.NewPostgresDB(os.Getenv("POSTGRES_URI"))
	tx := postgres.NewTxManager(postgresDB.DB)

	// REPOSITORIES
	repository := repository.NewRepository(postgresDB.DB)

	// OPENFGA
	openfgaClient, err := openfga.NewClient(
		os.Getenv("FGA_API_URL"),
		os.Getenv("FGA_STORE_ID"),
		os.Getenv("FGA_MODEL_ID"),
	)
	if err != nil {
		log.Fatalf("failed to create openfga client: %v", err)
	}

	// SERVICES
	authorizationSVC := authorization.NewService(openfgaClient)
	accountSVC := accounts.NewService(tx, repository)
	doctorSVC := doctors.NewService(tx, repository)
	nurseSVC := nurses.NewService(tx, repository)
	hospitalSVC := hospitals.NewService(tx, repository)
	medicalRecordSVC := medicalrecords.NewService(tx, repository, authorizationSVC)
	outboxSVC := outbox.NewService(repository)

	// RELAY
	r := relay.NewRelay(outboxSVC, authorizationSVC, 100, 2*time.Second)
	go r.Start(ctx)

	// SERVER
	server := http.NewServer(domain.APIName, domain.APIVersion, os.Getenv("PORT"), os.Getenv("API_AUTH_TOKEN"), accountSVC, doctorSVC, nurseSVC, hospitalSVC, medicalRecordSVC)
	server.Start()
}
