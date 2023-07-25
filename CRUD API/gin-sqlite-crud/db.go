// db.go
package main

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
    var err error
    db, err = sql.Open("sqlite3", "data.db")
    if err != nil {
        log.Fatal("Error opening database:", err)
    }

    createTableSQL := `
        CREATE TABLE IF NOT EXISTS items (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL
        );`

    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal("Error creating table:", err)
    }
}
