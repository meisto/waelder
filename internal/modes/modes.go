// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Wed 18 Jan 2023 02:19:46 AM CET
// Description: This package is mostly here to circumvent an input cycle 
// when modifying modi from subpages.
// ======================================================================
package modes

import (
   "github.com/muesli/termenv"

   ds "waelder/internal/datastructures"
)


type Mode int64
const (
   StartMode   Mode = iota
   ActiveMode  Mode = iota
   ActionMode  Mode = iota
   

   LoadingMode Mode = iota
   MainMode    Mode = iota
   ChoiceMode  Mode = iota
)

type ModeHandle struct {
   Update   func(*ds.Data, string)
   View     func(*termenv.Output, ds.Data, int, int, int, int)
}

var ModeLookup map[Mode]ModeHandle = map[Mode]ModeHandle {
   ActiveMode: {
      Update:  ActiveUpdate,
      View:    ActiveView,
   },
   ActionMode: actionModeHandle,
}

