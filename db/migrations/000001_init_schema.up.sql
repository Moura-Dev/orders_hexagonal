CREATE TABLE "companies"
(
    "id"         SERIAL PRIMARY KEY,
    "name"       VARCHAR     NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "users"
(
    "id"         SERIAL PRIMARY KEY,
    "company_id" SERIAL         NOT NULL,
    "name"       VARCHAR        NOT NULL,
    "email"      VARCHAR UNIQUE NOT NULL,
    "password"   VARCHAR        NOT NULL,
    "created_at" TIMESTAMPTZ    NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ    NOT NULL DEFAULT (now())
);

CREATE TABLE "user_companies"
(
    "id"         SERIAL PRIMARY KEY,
    "company_id" SERIAL,
    "user_id"    SERIAL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "roles"
(
    "id"         SERIAL PRIMARY KEY,
    "company_id" SERIAL,
    "user_id"    SERIAL,
    "name"       VARCHAR,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "user_roles"
(
    "id"         SERIAL PRIMARY KEY,
    "company_id" SERIAL,
    "user_id"    SERIAL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "permissions"
(
    "id"         SERIAL PRIMARY KEY,
    "company_id" SERIAL,
    "user_id"    SERIAL,
    "name"       VARCHAR,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "permission_roles"
(
    "id"         SERIAL PRIMARY KEY,
    "company_id" SERIAL,
    "user_id"    SERIAL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "contacts"
(
    "id"                 SERIAL PRIMARY KEY,
    "company_id"         SERIAL,
    "user_id"            SERIAL,
    "email"              VARCHAR,
    "website"            VARCHAR,
    "address"            VARCHAR,
    "inscricao_estadual" VARCHAR,
    "cnpj"               VARCHAR,
    "name"               VARCHAR,
    "cellphone"          VARCHAR,
    "logo_url"           VARCHAR,
    "fantasy_name"       VARCHAR,
    "created_at"         TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at"         TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "costumers"
(
    "id"         SERIAL PRIMARY KEY,
    "company_id" SERIAL,
    "user_id"    SERIAL,
    "contact_id" SERIAL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "factories"
(
    "id"         SERIAL PRIMARY KEY,
    "company_id" SERIAL,
    "contact_id" SERIAL,
    "user_id"    SERIAL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "catalogs"
(
    "id"         SERIAL PRIMARY KEY,
    "company_id" SERIAL,
    "user_id"    SERIAL,
    "factory_id" SERIAL,
    "data"       DATE,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "items"
(
    "id"          SERIAL PRIMARY KEY,
    "company_id"  SERIAL,
    "catalog_id"  SERIAL,
    "factory_id"  SERIAL,
    "user_id"     SERIAL,
    "code"        VARCHAR,
    "reference"   VARCHAR,
    "description" VARCHAR,
    "image_url"   VARCHAR,
    "price"       REAL,
    "ipi"         REAL,
    "discount"    SMALLINT,
    "quantity"    SMALLINT,
    "created_at"  TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at"  TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "orders"
(
    "id"          SERIAL PRIMARY KEY,
    "company_id"  SERIAL,
    "costumer_id" SERIAL,
    "contact_id"  SERIAL,
    "user_id"     SERIAL,
    "factory_id"  SERIAL,
    "sub_total"   REAL,
    "total"       REAL,
    "ipi"         REAL,
    "created_at"  TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at"  TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "order_items"
(
    "id"         SERIAL PRIMARY KEY,
    "company_id" SERIAL,
    "user_id"    SERIAL,
    "order_id"   SERIAL,
    "ipi"        REAL,
    "discount"   SMALLINT,
    "quantity"   SMALLINT,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "item_prices"
(
    "id"         SERIAL PRIMARY KEY,
    "company_id" SERIAL,
    "item_id"    SERIAL,
    "user_id"    SERIAL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "available_catalogs_by_company"
(
    "id"         SERIAL PRIMARY KEY,
    "company_id" SERIAL,
    "user_id"    SERIAL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "available_factories_by_company"
(
    "id"         SERIAL PRIMARY KEY,
    "company_id" SERIAL,
    "user_id"    SERIAL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions"
(
    "id"            UUID PRIMARY KEY,
    "email"         VARCHAR     NOT NULL,
    "refresh_token" VARCHAR     NOT NULL,
    "user_agent"    VARCHAR     NOT NULL,
    "client_ip"     VARCHAR     NOT NULL,
    "is_blocked"    BOOLEAN     NOT NULL DEFAULT false,
    "expires_at"    TIMESTAMPTZ NOT NULL,
    "created_at"    TIMESTAMPTZ NOT NULL DEFAULT (now())
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