.PHONY: dev dev-down test test-go test-openfga gen-mocks gen-openapi openfga-run openfga-stores-create openfga-models-create openfga-models-list seed-db openfga-init

dev: # [ make dev ] - spin up development stack
	@printf "\033[0;34mSpinning up Dev Environment...\033[0m\n"
	@echo
	@printf "\033[0;34mBuilding OpenHealth Services...\033[0m\n"
	@docker compose -f dev.docker-compose.yaml up --build -d --wait
	@echo
	@printf "\033[0;32mDev environment started!\033[0m\n"
	@echo
	@printf "\033[0;35m> Seed DB via [ make seed-db ] \033[0m\n"
	@echo
	@cd cmd/openhealth-app && bun --env-file=.env.development --bun run dev

dev-down: # [ make dev-down ] - teardown development stack
	docker compose -f dev.docker-compose.yaml down -v

API_AUTH_TOKEN ?= 12345
FGA_API_URL ?= http://localhost:8080
FGA_STORE_NAME ?= openhealth
export FGA_STORE_NAME

seed-db: # [ make seed-db ] - seed database with mock data
	@hurl --no-output --variable "API_AUTH_TOKEN=$(API_AUTH_TOKEN)" tools/hurl/seed.hurl
	@printf "\033[0;32m✨ Database seeded successfully\033[0m\n"

test: test-go test-openfga # [ make test ] - run ALL tests

test-go: # [ make test-go ] - run go tests
	@echo "\033[0;34m[ Go Tests ]\033[0m"
	@echo
	go test ./...

test-openfga: # [ make test-openfga ] - run openfga tests
	@echo "\033[0;35m[ OpenFGA Tests ]\033[0m"
	@echo
	fga model test --tests ${OPENFGA_TESTS}
	@echo "\033[0;32m✨ Success\033[0m"

gen-mocks: # [ make gen-mocks ] - generate mocks
	go generate ./...

gen-openapi: # [ make gen-openapi ] - generate openapi
	@go run ./scripts/gen-openapi/gen-openapi.go openapi > docs/openapi.yaml
	@echo "\033[0;34m> OpenAPI Generated\033[0m"
	@bash ./tools/bash/bruno-sync.sh

OPENFGA_TESTS=_openfga/openhealth.fga.yaml
OPENFGA_STORE_NAME ?= openhealth-openfga-store
OPENFGA_MODEL=_openfga/openhealth.fga

openfga-init: # [ make openfga-init ] - create or reuse the named store and ensure a model exists
	@docker compose -f dev.docker-compose.yaml run --build --rm openfga-init

openfga-run: # [ make openfga-run ] - run openfga container
	@docker compose -f dev.docker-compose.yaml up -d openfga > /dev/null

openfga-stores-create: # [ make openfga-stores-create NAME=store-name ] - generate a new openfga store
	@if [ -z "${OPENFGA_STORE_NAME}" ]; then echo "> Please provide a OPENFGA_STORE_NAME for the store" && exit 1; fi
	set -a && source .env && set +a && fga store create --name ${OPENFGA_STORE_NAME}

openfga-models-create: # [ make openfga-models-create ] - add a new openfga model to a store
	@STORE_ID=$$(FGA_API_URL="$(FGA_API_URL)" fga store list | jq -r --arg n "$(FGA_STORE_NAME)" '.stores[]? | select(.name == $$n) | .id' | head -n1); \
	if [ -z "$$STORE_ID" ]; then echo "OpenFGA store '$(FGA_STORE_NAME)' not found" >&2; exit 1; fi; \
	FGA_API_URL="$(FGA_API_URL)" FGA_STORE_ID="$$STORE_ID" fga model write --file "$(OPENFGA_MODEL)"

openfga-models-list: # [ make openfga-models-list ] - list all openfga models
	@STORE_ID=$$(FGA_API_URL="$(FGA_API_URL)" fga store list | jq -r --arg n "$(FGA_STORE_NAME)" '.stores[]? | select(.name == $$n) | .id' | head -n1); \
	if [ -z "$$STORE_ID" ]; then echo "OpenFGA store '$(FGA_STORE_NAME)' not found" >&2; exit 1; fi; \
	FGA_API_URL="$(FGA_API_URL)" FGA_STORE_ID="$$STORE_ID" fga model list
