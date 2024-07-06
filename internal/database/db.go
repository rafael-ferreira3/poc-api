package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	if err := Initialize(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
}

func Initialize() error {
	connStr := "user=postgres dbname=postgres password=apiserver sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		return err
	}

	log.Println("Database connection successfully initialized")
	return nil
}

func Close() {
	log.Println("Database connection closed")
	DB.Close()
}
