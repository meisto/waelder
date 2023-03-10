// ======================================================================
// Author: meisto
// Creation Date: Wed 18 Jan 2023 02:19:46 AM CET
// Description: This package is mostly here to circumvent an input cycle
// when modifying modi from subpages.
// ======================================================================
package modes

import (
	ds "waelder/internal/datastructures"
	"waelder/internal/renderer"
)

type Mode int64

const (
   NoMode      Mode = iota
	ActiveMode  Mode = iota
	ActionMode  Mode = iota
   MdViewMode  Mode = iota
   HelpMode    Mode = iota
)


var ModeLookup = map[Mode]func(ds.Data, int, int) renderer.RenderField{
	ActiveMode: activeView,
	ActionMode: actionView,
   MdViewMode: mdView,
   HelpMode: helpView,
}
