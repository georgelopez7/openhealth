.PHONY: test test-go test-openfga gen-mocks gen-openapi run-openfga gen-openfga-store add-openfga-model list-openfga-models seed-db

dev: # [ make dev ]
	@printf "\033[0;34mSpinning up Dev Environment...\033[0m\n"
	@echo
	@make init-openfga PRINT_ENV=false
	@echo
	@printf "\033[0;34mBuilding OpenHealth Services...\033[0m\n"
	@docker compose -f dev.docker-compose.yaml up --build -d > /dev/null
	@echo
	@printf "\033[0;32mDev environment started!\033[0m\n"
	@echo
	@printf "\033[0;35m> Seed DB via [ make seed-db ] \033[0m\n"
	@echo

# Seed the database with mock data
seed-db: # [ make seed-db ]
	@hurl --no-output _hurl/seed.hurl
	@printf "\033[0;32m✨ Database seeded successfully\033[0m\n"

# Teardown dev containers and remove volumes
dev-down: # [ make dev-down ]
	docker compose -f dev.docker-compose.yaml down -v

# Run All Tests
test: test-go test-openfga # [ make test ]

# Run Go Tests
test-go: # [ make test-go ]
	@echo "\033[0;34m[ Go Tests ]\033[0m"
	@echo
	go test ./...

# Run OpenFGA Tests
test-openfga: # [ make test-openfga ]
	@echo "\033[0;35m[ OpenFGA Tests ]\033[0m"
	@echo
	fga model test --tests ${OPENFGA_TESTS}
	@echo "\033[0;32m✨ Success\033[0m"

# Generate Mocks
gen-mocks: # [ make gen-mocks ]
	go generate ./...

# Generate OpenAPI
gen-openapi: # [ make gen-openapi ]
	@go run ./scripts/gen-openapi/gen-openapi.go openapi > docs/openapi.yaml
	@echo "\033[0;34m> OpenAPI Generated\033[0m"
	@bash ./_bash/bruno-sync.sh

################
## OPENFGA ##
################

OPENFGA_TESTS=_openfga/openhealth.fga.yaml
OPENFGA_STORE_NAME ?= openhealth-openfga-store
OPENFGA_MODEL=_openfga/openhealth.fga

PRINT_ENV ?= true

# Init OpenFGA Store & Model & Update Environment Variables
init-openfga: # [ make init-openfga ]
	@make run-openfga
	@touch .env && \
	set -a && source .env && set +a && \
	unset FGA_STORE_ID && \
	STORE=$$(fga store create --name openhealth) && \
	STORE_ID=$$(echo $$STORE | jq -r '.store.id') && \
	export FGA_STORE_ID=$$STORE_ID && \
	echo "" && \
	printf "\033[0;32m✓ OpenFGA Store Created\033[0m\n" && \
	MODEL=$$(fga model write --file $(OPENFGA_MODEL)) && \
	MODEL_ID=$$(echo $$MODEL | jq -r '.authorization_model_id') && \
	printf "\033[0;32m✓ OpenFGA Model Created & Added To Store\033[0m\n" && \
	if [ "$(PRINT_ENV)" != "false" ]; then \
		echo "---" && \
		echo "FGA_API_URL=http://localhost:8080" && \
		echo "FGA_STORE_ID=$$STORE_ID" && \
		echo "FGA_MODEL_ID=$$MODEL_ID" && \
		echo "---"; \
	fi && \
	for pair in \
		"FGA_API_URL=http://localhost:8080" \
		"FGA_STORE_ID=$$STORE_ID" \
		"FGA_MODEL_ID=$$MODEL_ID"; do \
		key=$${pair%%=*}; \
		grep -vE "^$${key}=" .env > .env.tmp && mv .env.tmp .env; \
		echo "$$pair" >> .env; \
	done && \
	printf "\033[0;32m✓ Environment Variables Updated\033[0m\n"


# Run OpenFGA Container
run-openfga: # [ make run-openfga ]
	@docker compose -f dev.docker-compose.yaml up -d openfga > /dev/null

# Generate a new OpenFGA store
gen-openfga-store: # [ make gen-openfga-store NAME=store-name ]
	@if [ -z "${OPENFGA_STORE_NAME}" ]; then echo "> Please provide a OPENFGA_STORE_NAME for the store" && exit 1; fi
	set -a && source .env && set +a && fga store create --name ${OPENFGA_STORE_NAME}

# Add a new OpenFGA model to a store
add-openfga-model: # [ make add-openfga-model ]
	set -a && source .env && set +a && fga model write --file ${OPENFGA_MODEL}

# List all OpenFGA models
list-openfga-models: # [ make list-openfga-models ]
	set -a && source .env && set +a && fga model list