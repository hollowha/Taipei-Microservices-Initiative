package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DBConn() *sql.DB {
	connStr := "host=localhost user=yourusername password=yourpassword dbname=yourdbname sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	log.Println("Connected to the database successfully")
	return db
}
