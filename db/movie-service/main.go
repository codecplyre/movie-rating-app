package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/")
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}
	// Create the database if it doesn't exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS movie_service")
	if err != nil {
		panic(err.Error())
	}
	// Use the database
	_, err = db.Exec("USE movie_service")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("MySQL database created and selected successfully")
	// run migrations
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("failed to get database driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql", driver)
	if err != nil {
		log.Fatalf("failed to create migration instance: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to apply migration: %v", err)
	}

	fmt.Println("migration successful")
}
