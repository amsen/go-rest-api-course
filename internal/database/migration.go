package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// MigrateDB - runs all migrations in the migrations
func (db *Database) MigrateDB() error {
	fmt.Println("Running migrations for the database")

	driver, err := postgres.WithInstance(db.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create the postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver)

	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := m.Up(); err != nil {
		fmt.Println("Migrations did not work properly...", err)
	}

	fmt.Println("Completed migrations for the database")

	return nil
}
