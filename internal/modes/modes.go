// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Wed 18 Jan 2023 02:19:46 AM CET
// Description: This package is mostly here to circumvent an input cycle 
// when modifying modi from subpages.
// ======================================================================
package modes

import (
   ds "waelder/internal/datastructures"

   tea "github.com/charmbracelet/bubbletea"
)


type Mode int64
const (
   StartMode   Mode = iota
   LoadingMode Mode = iota
   MainMode    Mode = iota
   ActiveMode  Mode = iota
   ChoiceMode  Mode = iota
)

type ModeHandle struct {
   Update   func(*ds.Data, tea.KeyMsg)
   View     func(ds.Data, int, int) []string
}


/* 
var ModeLookup map[Mode]ModeHandle = map[Mode]ModeHandle {
   ActiveMode: {
      Update:  ActiveUpdate,
      View:    ActiveView,
   },
}

*/

