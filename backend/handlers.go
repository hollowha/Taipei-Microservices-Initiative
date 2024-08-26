package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	db := DBConn() // 調用db.go中的DBConn()來連接資料庫
	defer db.Close()

	rows, err := db.Query("SELECT name FROM users")
	if err != nil {
		log.Fatal("Error querying the database: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal("Error scanning the row: ", err)
		}
		w.Write([]byte("User: " + name + "\n"))
	}
}
