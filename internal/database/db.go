package database

import "database/sql"

var DB *sql.DB

func Initialize() error {
	connStr := "user=postgres dbname=postgres password=apiserver sslmode=disable"
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		return err
	}

	return nil
}

func close() {
	DB.Close()
}
