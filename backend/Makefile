include .env

MIGRATION_PATH = ./migrations/migration.sql

db-migrate:
	psql -f $(MIGRATION_PATH) $(POSTGRES_URL)

test: db-migrate
	go test ./...