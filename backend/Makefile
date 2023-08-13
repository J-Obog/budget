include .env

MIGRATION_PATH = ./migrations/migration.sql
QUEUE_NAMES = test.queue category.deleted

db-migrate:
	psql -f $(MIGRATION_PATH) $(POSTGRES_URL)

create-queues:
	for q in $(QUEUE_NAMES); do \
		python rabbitmqadmin declare queue name=$$q -U $(RABBIT_MQ_ADMIN_URL); \
	done

test: db-migrate create-queues
	export APP_ENV=dev && gotest ./... -count=1 -v