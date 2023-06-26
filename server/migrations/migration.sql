CREATE TABLE IF NOT EXISTS accounts (
    "id" VARCHAR PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL,
	"password" VARCHAR NOT NULL,
	"is_activated" BOOLEAN NOT NULL,
	"created_at" BIGINT NOT NULL,
	"updated_at" BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS budgets (
    "id" VARCHAR PRIMARY KEY,
	"account_id" VARCHAR NOT NULL,
	"category_id" VARCHAR NULL,
	"name" VARCHAR NOT NULL,
	"type" SMALLINT NOT NULL,
	"month" INTEGER NOT NULL,
	"year" INTEGER NOT NULL,
	"projected" FLOAT NOT NULL,
	"actual" FLOAT NOT NULL,
	"created_at" BIGINT NOT NULL,
	"updated_at" BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
    "id" VARCHAR PRIMARY KEY,
	"account_id" VARCHAR NOT NULL,
	"budget_id" VARCHAR NOT NULL,
	"description" VARCHAR NULL,
	"amount" FLOAT NOT NULL,
	"month" INTEGER NOT NULL,
	"year" INTEGER NOT NULL,
	"day" INTEGER NOT NULL,
	"created_at" BIGINT NOT NULL,
	"updated_at" BIGINT NOT NULL
)