// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 30 Jan 2023 03:11:24 PM CET
// Description: -
// ======================================================================
package modes

import (
   "fmt"
   "sort"

   "github.com/muesli/termenv"

   ds "waelder/internal/datastructures"
   "waelder/internal/renderer"
)

func helpView(
   output *termenv.Output,
   d ds.Data,
   windowHeight int,
   windowWidth int,
) renderer.RenderField {
   dist := renderer.GenerateNode(" | ", "dead")
   
   var s []renderer.Renderable = []renderer.Renderable{
      renderer.GenerateNoRenderNode(" "),
      renderer.GenerateNode("q: quit", "dead"),
      dist,
   }

   // Generate all nodes
   edges := d.Graph.GetEdges(d.Graph.ActiveNode)
   sort.Slice(edges, func(i, j int) bool{
      return edges[i].Label < edges[j].Label
   })

   l := 0
   for _, i := range(s) { l += i.Length() }

   for i := 0; i < len(edges); i++ {
      msg := fmt.Sprintf("%s: %s", edges[i].Label, edges[i].Description)
      x := renderer.GenerateNode(msg, "dead")
   
      if l + x.Length() < windowWidth - 3 {
         s = append(s, x)
         if i < len(edges) - 1 { s = append(s, dist) }
      } else {
         if i < len(s) - 1 {
            s = append(s, renderer.GenerateNode("...", "dead"))
         }
         break
      }

      l += x.Length() + dist.Length()

   }

   return renderer.GenerateField([]renderer.RenderLine{
      renderer.GenerateLine(windowWidth, s),
   })
}
