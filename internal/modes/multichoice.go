// ======================================================================
// Author: meisto
// Creation Date: Wed 25 Jan 2023 04:56:19 PM CET
// Description: -
// ======================================================================
package modes

import(
   "waelder/internal/renderer"
)

type MultiChoice struct {
   index    int
   selected []int
   header   string
   footer   string
   buttons  []string
   labels   []string
}
func GetMultiChoice(
   index int,
   selected []int,
   header, 
   footer string,
   buttons,
   labels []string,
) MultiChoice {
   return MultiChoice{index, selected, header, footer, buttons, labels}
}
func (c *MultiChoice) Step() {
   c.index += 1
   if c.index > len(c.buttons) { c.index = 0}
}
func (c MultiChoice) ToRenderField(x,y, width int) renderer.RenderField{

   fActive := func(c string) renderer.Renderable {
      return renderer.GenerateNode(c, "active")
   }
   f := func(c string) renderer.Renderable { return renderer.GenerateNoRenderNode(c)}


   rl := make([]renderer.RenderLine, len(c.buttons) + 1)
   rl[0] = renderer.GenerateLine(
      width - 2,
      []renderer.Renderable{renderer.GenerateNoRenderNode(c.header)},
   )
   for i := 0; i < len(c.buttons); i++{
      var c1 renderer.Renderable = f(c.buttons[i])
      var c2 renderer.Renderable = f(" " + c.labels[i])

      for _, j := range(c.selected) {
         if i == c.index || i == j{
            c1 = fActive(c.buttons[i])
            c2 = f(" " + c.labels[i])
         } 
      }

      rl[i + 1] = renderer.GenerateLine(width-2, []renderer.Renderable{c1, c2})
   }

   return renderer.GenerateField(rl)
}
func (c MultiChoice) GetSelection() []string { 
   var s []string

   for _, j := range(c.selected) { s = append(s, c.labels[j]) }
   return s
}
