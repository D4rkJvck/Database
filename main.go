package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Initialize Database
	db, err_db := sql.Open("sqlite3", "./database/db.sqlite")
	if err_db != nil {
		log.Fatalln("ERROR: Open Database Failed:\n", err_db)
	}

	defer db.Close()

	// Open Queries File
	queries, err_sql := os.ReadFile("./database/tables.sql")
	if err_sql != nil {
		// TODO: SQL Queries Models
		log.Fatalln("ERROR: Read File Failed:\n", err_sql)
	}

	// Create Tables
	_, err_exe := db.Exec(string(queries))
	if err_exe != nil {
		log.Fatalln("ERROR: Execute Query Failed:\n", err_exe)
	}
}
