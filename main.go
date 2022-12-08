package main

import (
	"database/sql"
	"orders_hexagonal/db"
	"orders_hexagonal/server"
	"orders_hexagonal/util"
)

func main() {
	config, _ := util.LoadConfig(".")
	conn, _ := sql.Open(config.DBDriver, config.DBSource)
	driver, _ := postgres.WithInstance(conn, &postgres.Config{})

	runDBMigration(config.MigrationURL, config.DBDriver, driver)

	store := db.NewStorage(conn)

	runGinServer(config, store)
}

func runDBMigration(migrationURL string, dbSource string, driver database.Driver) {
	m, _ := migrate.NewWithDatabaseInstance(migrationURL, dbSource, driver)

	_ = m.Up()
}

func runGinServer(config util.Config, store db.Store) {
	api := server.NewServer(config, store)

	_ = api.Start(config.HTTPServerAddress)
}
