// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Wed 16 Nov 2022 11:18:35 PM CET
// Description: -
// ======================================================================
package main

import (
   "fmt"
   "os"
   tea "github.com/charmbracelet/bubbletea"

   m "dntui/internal/model"
   "dntui/internal/io"
)


func main() {
   
   // Load database
   dbHandle := io.GetDatabaseHandle()
   defer dbHandle.Close()

   q := "CREATE TABLE IF NOT EXISTS tab(x INTEGER);"
   dbHandle.Exec(q)
   dbHandle.Close()

   os.Exit(0)

   p := tea.NewProgram(m.InitialModel(dbHandle))
   if _, err := p.Run(); err != nil {
      fmt.Printf("Error: %v", err)
      os.Exit(1)
   }
}
