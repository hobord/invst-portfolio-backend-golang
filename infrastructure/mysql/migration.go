package mysql_repository

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4/database/mysql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrationUp(conn *sql.DB, migrationFilesPath string) {
	driver, err := mysql.WithInstance(conn, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationFilesPath,
		"mysql", driver)
	if err != nil {
		log.Fatal(err)
	}
	// m.Steps(2)
	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}
}
func MigrationDown(conn *sql.DB, migrationFilesPath string) {
	driver, err := mysql.WithInstance(conn, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationFilesPath,
		"mysql", driver)

	if err != nil {
		log.Fatal(err)
	}
	err = m.Down()
	if err != nil {
		log.Fatal(err)
	}
}
