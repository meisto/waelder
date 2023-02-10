// ======================================================================
// Author: meisto
// Creation Date: Wed 18 Jan 2023 03:05:03 PM CET
// Description: -
// ======================================================================
package modes

import (
	"github.com/muesli/termenv"

	ds "waelder/internal/datastructures"
	"waelder/internal/renderer"
)

func actionView(output *termenv.Output, d ds.Data, height int, width int) renderer.RenderField {

   var res []renderer.RenderLine

   for i := len(d.CombatLog.Current.Actions) - 1; i >= 0 && height > 0; i-- {
      e := d.CombatLog.Current.Actions[i]
      res = append(res, renderer.GenerateLineFromOne(width, e.Display(d)))
      height -= 1
   }

   for i := len(d.CombatLog.PreviousRounds) - 1; i >= 0 && height > 0; i -- {
      x := d.CombatLog.PreviousRounds[i]
      for j := len(x.Actions) - 1; j >= 0 && height > 0; j -- {
         e := x.Actions[j]
         res = append(res, renderer.GenerateLineFromOne(width, e.Display(d)))
         height -= 1
      }
   }

   res2 := []renderer.RenderLine{}
   for i := len(res) - 1; i >= 0; i-- {
      res2 = append(res2, res[i])
   }

   return renderer.GenerateField(res2)
}
