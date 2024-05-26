CREATE TABLE "users"
(
    "id"         bigserial PRIMARY KEY,
    "username"   varchar UNIQUE NOT NULL,
    "name"       varchar        NOT NULL,
    "email"      varchar        NOT NULL,
    "password"   varchar        NOT NULL,
    "created_at" timestamptz    NOT NULL DEFAULT (now())
);