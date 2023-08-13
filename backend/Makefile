include .env

MIGRATION_PATH = ./migrations/migration.sql

db-migrate:
	psql -f $(MIGRATION_PATH) $(POSTGRES_URL)

create-queues:
	python rabbitmqadmin declare queue name="test.queue" -U $(RABBIT_MQ_ADMIN_URL)
	python rabbitmqadmin declare queue name="category.deleted" -U $(RABBIT_MQ_ADMIN_URL)

test: db-migrate create-queues
	export APP_ENV=dev && gotest ./... -count=1 -v