// ======================================================================
// Author: meisto
// Creation Date: Thu 26 Jan 2023 04:39:22 PM CET
// Description: -
// ======================================================================
package wio

import (
   "log"
   "os"
   "path"
)

func ReadLocalFileToString(filepath string) string {
   root, _ := os.Getwd()
   fp := path.Join(root, filepath)

   content, err := os.ReadFile(fp)
   if err != nil {
      log.Fatal("File not found '", fp, "'.")
   }

   return string(content)
}

