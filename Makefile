
include .env
MIGRATIONS_DIR    = ./cmd/migrate/migrations
MIGRATIONS_PATH = file://${MIGRATIONS_DIR}

.PHONY: migrate-create
migrate-create:
	@migrate create -seq -ext sql -dir ${MIGRATIONS_DIR} $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up: 
	@migrate -verbose -source="${MIGRATIONS_PATH}" -database="postgres://postgres:postgres@localhost:5434/go-social?sslmode=disable" up

.PHONY: migrate-down
migrate-down: 
	@migrate -verbose -source="${MIGRATIONS_PATH}" -database="postgres://postgres:postgres@localhost:5434/go-social?sslmode=disable" down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: seed
seed:
	@go run cmd/migrate/seed/main.go