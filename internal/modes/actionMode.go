// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Wed 18 Jan 2023 03:05:03 PM CET
// Description: -
// ======================================================================
package modes

import(
   "github.com/muesli/termenv"

   ds "waelder/internal/datastructures"
   "waelder/internal/renderer"
)

var actionModeHandle ModeHandle = ModeHandle {
   Update: func(*ds.Data, string) {},
   View: func(output *termenv.Output, d ds.Data, x int, y int, height int, width int) { 

      if len(d.CombatLog.Actions) >  0 {

         var res []renderer.RenderLine
         for _, e := range(d.CombatLog.Actions) {
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

      renderer.GenerateField(res).RenderBlock(output, x, y, height, true, true)

      }


   },
}
