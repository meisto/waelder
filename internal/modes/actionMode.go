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
   if len(d.CombatLog.Current.Actions) > 0 {

      for _, e := range d.CombatLog.Current.Actions {
         res = append(
            res,
            renderer.GenerateLine(
               width,
               []renderer.Renderable{
                  renderer.GenerateNode(e.Display(), "default"),
               },
            ),
         )
      }
   }
   return renderer.GenerateField(res)
}
