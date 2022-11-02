CREATE SCHEMA "sample_bank";

CREATE TABLE "sample_bank"."accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sample_bank"."entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sample_bank"."transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "sample_bank"."accounts" ("owner");

CREATE INDEX ON "sample_bank"."entries" ("account_id");

CREATE INDEX ON "sample_bank"."transfers" ("from_account_id");

CREATE INDEX ON "sample_bank"."transfers" ("to_account_id");

CREATE INDEX ON "sample_bank"."transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "sample_bank"."entries"."amount" IS 'positive  or negative number';

COMMENT ON COLUMN "sample_bank"."transfers"."amount" IS 'must be positive number';

ALTER TABLE "sample_bank"."entries" ADD FOREIGN KEY ("account_id") REFERENCES "sample_bank"."accounts" ("id");

ALTER TABLE "sample_bank"."transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "sample_bank"."accounts" ("id");

ALTER TABLE "sample_bank"."transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "sample_bank"."accounts" ("id");
