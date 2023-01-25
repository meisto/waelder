// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 16 Jan 2023 11:23:52 PM CET
// Description: -
// ======================================================================
package queries

const (
   InitializeCharacterTable string = `
      CREATE TABLE IF NOT EXISTS character(
         name TEXT PRIMARY KEY,
         affiliation TEXT,
         race TEXT NOT NULL,
         class TEXT
      );
   `

   InitializeGenericEnemyTable string = `
      CREATE TABLE IF NOT EXISTS genericEnemy(
         name TEXT PRIMARY KEY,
         shortDesc TEXT NOT NULL,
         race TEXT,
         alignment INTEGER,
         descPath TEXT
      );
   `

   FetchCharacterNames string = `
      SELECT name FROM character;
   `

   AddCharacter string = `
      INSERT INTO character(name, affiliation, race, class)
         VALUES (?, ?, ?, ?);
   `

   GetCharacterByName string = "SELECT * FROM character WHERE name == ?;"




)

