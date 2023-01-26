// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Wed 18 Jan 2023 03:05:03 PM CET
// Description: -
// ======================================================================
package modes

import (
	"github.com/muesli/termenv"

	ds "waelder/internal/datastructures"
	"waelder/internal/renderer"
)

func actionView(output *termenv.Output, d ds.Data, x int, y int, height int, width int) {

   if len(d.CombatLog.Current.Actions) > 0 {

      var res []renderer.RenderLine
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

      renderer.GenerateField(res).RenderBlock(output, x, y, height, false, 1000)  

   }
}
