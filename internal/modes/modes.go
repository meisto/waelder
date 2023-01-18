// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Wed 18 Jan 2023 02:19:46 AM CET
// Description: This package is mostly here to circumvent an input cycle 
// when modifying modi from subpages.
// ======================================================================
package modi

type Mode int64
const (
   StartMode   Mode = iota
   LoadingMode Mode = iota
   MainMode    Mode = iota
   ActiveMode  Mode = iota
   ChoiceMode  Mode = iota
)



