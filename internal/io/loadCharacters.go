// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 16 Jan 2023 01:20:07 AM CET
// Description: -
// ======================================================================

package io

import (
   "encoding/json"
   "fmt"
   "io/ioutil"
   "log"

   cm "dntui/internal/datastructures"
)

func LoadCharacterFromFile(filePath string) cm.Character {

   data, err := ioutil.ReadFile(filePath)
   if err != nil {
      log.Print(fmt.Sprintf("[WARNING] Could not find file '%s'.", filePath))
      log.Fatal(err)
   }

   var char cm.Character

   err = json.Unmarshal(data, &char)
   if err != nil {log.Fatal(err)}


   return char
}

// func loadParty()

