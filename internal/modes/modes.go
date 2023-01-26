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
   MdViewMode  Mode = iota
)


var ModeLookup = map[Mode]func(*termenv.Output, ds.Data, int, int, int, int) {
	ActiveMode: activeView,
	ActionMode: actionView,
   MdViewMode: mdView,
}
