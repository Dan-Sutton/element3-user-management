package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
    var err error
    connStr := "user=myuser password=mypassword dbname=mydatabase sslmode=disable"
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    createTable()
}

func createTable() {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        first_name VARCHAR(100),
        surname VARCHAR(100),
        email VARCHAR(100),
        dob DATE
    );`
    _, err := DB.Exec(query)
    if err != nil {
        log.Fatal(err)
    }
}