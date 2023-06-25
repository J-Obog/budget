CREATE TABLE IF NOT EXISTS accounts (
    "id" VARCHAR PRIMARY KEY,
    "name" VARCHAR,
    "email" VARCHAR,
	"password" VARCHAR,
	"is_activated" BOOLEAN,
	"created_at"   BIGINT,
	"updated_at" BIGINT
)