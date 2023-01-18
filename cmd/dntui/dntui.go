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

   root "dntui/internal"
   "dntui/internal/io"
)


func main() {
   
   // Load database
   dbHandle := io.GetDatabaseHandle()
   defer dbHandle.Close()

   io.CreateTables(dbHandle)

   p := tea.NewProgram(root.InitialModel(dbHandle))
   if _, err := p.Run(); err != nil {
      fmt.Printf("Error: %v", err)
      os.Exit(1)
   }
}
