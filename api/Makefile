include .env
MIGRATIONS_PATH = cmd/migrations

.PHONY test:
test:
	@go test ./...

.PHONY run:
run:
	@go run cmd/main.go

.PHONY: migrate-create
migrate-create:
	@name=$(name);
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(name)

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(POSTGRES_URL) up

.PHONY: migrate-down
migrate-down:
	@name=$(name);
	@migrate -path=$(MIGRATIONS_PATH) -database=$(POSTGRES_URL) down $(name)