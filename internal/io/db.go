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

   "waelder/internal/io/queries"
   cm "waelder/internal/datastructures"
)

func GetDatabaseHandle() *sql.DB {
   dbPath := os.Getenv("waelder_db_path")
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
   fmt.Printf("Successfully connected to SQLite database at '%s'.\n", dbPath)

   return db
}

func CreateTables(handle *sql.DB) {
   handle.Exec(queries.InitializeCharacterTable)

   _, err := handle.Exec(queries.InitializeGenericEnemyTable)
   if err != nil {
      log.Fatal("Error when initializing generic enemy table")
   }
}


func ReadCharacterFromDatabase(handle *sql.DB, charName string) cm.Character {

   row := handle.QueryRow(queries.GetCharacterByName, charName)
   
   var name          string
   var affiliation   string
   var race          string
   var class         string

   err := row.Scan(&name, &affiliation, &race, &class)
   if err != nil { log.Fatal(err)}
   
   return cm.Character {
      Name:          name,
      Affiliation:   affiliation,
      Race:           race,
      // Subrace     string
      Class:         class,
      Stats: cm.CharacterStats {
         Hp: 10,
         Max_hp: 10,
         Initiative: 10,
      },
   }
}

/**
   Synchronize a characte with the database. If the Character does not exist in
   the database add it.
   If it does exist the database is assumed to be up-to-date and the 
   Character struct is Updated.
**/
func SyncCharacterWithDatabase(handle *sql.DB, char cm.Character) cm.Character {

   rows, err := handle.Query(queries.FetchCharacterNames)
   if err != nil { log.Fatal(err) }
   defer rows.Close()

   for rows.Next() {
      var readName   string

      rows.Scan(&readName)
     
      if readName == char.Name {
         log.Print(fmt.Sprintf("Sync character '%s' with database.", char.Name))

         return ReadCharacterFromDatabase(handle, char.Name)
      }
   }


   {
      log.Print(fmt.Sprintf("Add character '%s' to database.", char.Name))
      _, err := handle.Exec(
         queries.AddCharacter, 
         char.Name,
         char.Affiliation,
         char.Race,
         char.Class,
      )
      if err != nil {log.Fatal(err)}
   }

   return char
}
