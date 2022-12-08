CREATE TABLE "companies"
(
    "id"         serial PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users"
(
    "id"         serial PRIMARY KEY,
    "company_id" serial,
    "name"       varchar        NOT NULL,
    "email"      varchar UNIQUE NOT NULL,
    "password"   varchar        NOT NULL,
    "created_at" timestamptz    NOT NULL DEFAULT (now()),
    "updated_at" timestamptz    NOT NULL DEFAULT (now())
);

CREATE TABLE "user_companies"
(
    "id"         serial PRIMARY KEY,
    "company_id" serial,
    "user_id"    serial,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "roles"
(
    "id"         serial PRIMARY KEY,
    "company_id" serial,
    "user_id"    serial,
    "name"       varchar,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_roles"
(
    "id"         serial PRIMARY KEY,
    "company_id" serial,
    "user_id"    serial,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "permissions"
(
    "id"         serial PRIMARY KEY,
    "company_id" serial,
    "user_id"    serial,
    "name"       varchar,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "permission_roles"
(
    "id"         serial PRIMARY KEY,
    "company_id" serial,
    "user_id"    serial,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "contacts"
(
    "id"         serial PRIMARY KEY,
    "company_id" serial,
    "user_id"    serial,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "costumers"
(
    "id"         serial PRIMARY KEY,
    "company_id" serial,
    "user_id"    serial,
    "contact_id" serial,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "factories"
(
    "id"         serial PRIMARY KEY,
    "company_id" serial,
    "contact_id" serial,
    "user_id"    serial,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "catalogs"
(
    "id"         serial PRIMARY KEY,
    "company_id" serial,
    "user_id"    serial,
    "factory_id" serial,
    "data"       date,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "items"
(
    "id"          serial PRIMARY KEY,
    "company_id"  serial,
    "catalog_id"  serial,
    "factory_id"  serial,
    "user_id"     serial,
    "code"        varchar,
    "reference"   varchar,
    "description" text,
    "image_url"   text,
    "price"       real,
    "ipi"         real,
    "discount"    smallint,
    "quantity"    smallint,
    "created_at"  timestamptz NOT NULL DEFAULT (now()),
    "updated_at"  timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orders"
(
    "id"          serial PRIMARY KEY,
    "company_id"  serial,
    "costumer_id" serial,
    "contact_id"  serial,
    "user_id"     serial,
    "factory_id"  serial,
    "sub_total"   real,
    "total"       real,
    "ipi"         real,
    "created_at"  timestamptz NOT NULL DEFAULT (now()),
    "updated_at"  timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "order_items"
(
    "id"         serial PRIMARY KEY,
    "company_id" serial,
    "user_id"    serial,
    "order_id"   serial,
    "ipi"        real,
    "discount"   smallint,
    "quantity"   smallint,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "item_prices"
(
    "id"         serial PRIMARY KEY,
    "company_id" serial,
    "item_id"    serial,
    "user_id"    serial,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "available_catalogs_by_company"
(
    "id"         serial PRIMARY KEY,
    "company_id" serial,
    "user_id"    serial,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "available_factories_by_company"
(
    "id"         serial PRIMARY KEY,
    "company_id" serial,
    "user_id"    serial,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions"
(
    "id"            uuid PRIMARY KEY,
    "email"         varchar     NOT NULL,
    "refresh_token" varchar     NOT NULL,
    "user_agent"    varchar     NOT NULL,
    "client_ip"     varchar     NOT NULL,
    "is_blocked"    boolean     NOT NULL DEFAULT false,
    "expires_at"    timestamptz NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "users"
    ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "user_companies"
    ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "user_companies"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "roles"
    ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "roles"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "user_roles"
    ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "user_roles"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "permissions"
    ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "permissions"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "permission_roles"
    ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "permission_roles"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "contacts"
    ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "contacts"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "costumers"
    ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "costumers"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "costumers"
    ADD FOREIGN KEY ("contact_id") REFERENCES "contacts" ("id") ON DELETE CASCADE;

ALTER TABLE "factories"
    ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "factories"
    ADD FOREIGN KEY ("contact_id") REFERENCES "factories" ("id") ON DELETE CASCADE;

ALTER TABLE "factories"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "catalogs"
    ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "catalogs"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "catalogs"
    ADD FOREIGN KEY ("factory_id") REFERENCES "factories" ("id") ON DELETE CASCADE;

ALTER TABLE "items"
    ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "items"
    ADD FOREIGN KEY ("catalog_id") REFERENCES "catalogs" ("id") ON DELETE CASCADE;

ALTER TABLE "items"
    ADD FOREIGN KEY ("factory_id") REFERENCES "factories" ("id") ON DELETE CASCADE;

ALTER TABLE "items"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "orders"
    ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "orders"
    ADD FOREIGN KEY ("costumer_id") REFERENCES "costumers" ("id") ON DELETE CASCADE;

ALTER TABLE "orders"
    ADD FOREIGN KEY ("contact_id") REFERENCES "costumers" ("id") ON DELETE CASCADE;

ALTER TABLE "orders"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "orders"
    ADD FOREIGN KEY ("factory_id") REFERENCES "factories" ("id") ON DELETE CASCADE;

ALTER TABLE "order_items"
    ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "order_items"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "order_items"
    ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id") ON DELETE CASCADE;

ALTER TABLE "item_prices"
    ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "item_prices"
    ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id") ON DELETE CASCADE;

ALTER TABLE "item_prices"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "available_catalogs_by_company"
    ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "available_catalogs_by_company"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "available_factories_by_company"
    ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "available_factories_by_company"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "sessions"
    ADD FOREIGN KEY ("email") REFERENCES "users" ("email") ON DELETE CASCADE;