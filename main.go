package main

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"project-orders/db"
	"project-orders/server"
	"project-orders/util"
)

func main() {
	config, _ := util.LoadConfig(".")
	conn, _ := sql.Open(config.DBDriver, config.DBSource)
	driver, _ := postgres.WithInstance(conn, &postgres.Config{})

	runDBMigration(config.MigrationURL, config.DBDriver, driver)

	storage := db.NewStore(conn)

	runGinServer(config, storage)
}

func runDBMigration(migrationURL string, dbSource string, driver database.Driver) {
	m, _ := migrate.NewWithDatabaseInstance(migrationURL, dbSource, driver)

	_ = m.Up()
}

func runGinServer(config util.Config, storage db.Storage) {
	api := server.NewGin(config, storage)

	_ = api.Start(config.HTTPServerAddress)
}
