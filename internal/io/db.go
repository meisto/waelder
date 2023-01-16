// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 16 Jan 2023 07:20:32 PM CET
// Description: -
// ======================================================================
package io

import (
   "fmt"
   "os"
   "log"

   "database/sql"
   _ "github.com/mattn/go-sqlite3"
)

func GetDatabaseHandle() *sql.DB {
   dbPath := os.Getenv("dntui_db_path")
   if dbPath == "" {
      log.Fatal("Could not find database path in environment")
   }

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

   // Test ping the db
   errPing := db.Ping()
   if errPing != nil {
      log.Fatal(errPing)
   }
   fmt.Printf("Successfully connected to SQLite database at '%s'", dbPath)

   return db
}

