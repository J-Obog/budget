

MIGRATIONS_PATH = ./migrations/migration.sql
POSTGRES_URL = $(POSTGRES_URL)
PGDB = paidoff 
PGHOST = localhost

foo:
	psql -d paidoff -U postgres -f $(MIGRATIONS_PATH)

test:
	go test ./...