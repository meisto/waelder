// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 16 Jan 2023 01:20:07 AM CET
// Description: -
// ======================================================================

package io

import (
   // "encoding/json"
   "fmt"
   "os"
)

func GetFiles(dirPath string) []string {

   _, err := os.Stat(dirPath)

   if os.IsNotExist(err) {
      fmt.Fprintf(os.Stderr, "Error when trying to open file '" + dirPath + "'")
      return []string{}
   }

   f, err := os.Open(dirPath)

   files, err := f.ReadDir(0)
   
   /*
   readFile, err := os.Open(filePath)
   if err != nil {
      fmt.Fprintf(os.Stderr, "Error when trying to open file '" + filePath + "'")
      return []string{}
   }

   readFile.Chdir
   */

   res := []string{}
   for _, x := range(files) {
      res = append(res, x.Name())
   } 


   return res

}

// func loadParty()

