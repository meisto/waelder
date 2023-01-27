// ======================================================================
// Author: meisto
// Creation Date: Tue 17 Jan 2023 05:04:20 PM CET
// Description: -
// ======================================================================
package modes

import (
   "waelder/internal/renderer"
)

type Choice struct {
   index    int
   Header   string
   Footer   string
   buttons  []string
   labels   []string
}
func GetChoice(index int, header, footer string, buttons []string, labels []string) Choice {
   return Choice{index, header, footer, buttons, labels}
}
func (c *Choice) Step() {
   c.index += 1
   if c.index >= len(c.buttons) { c.index = 0}
}

func (c *Choice) GoTo(n int) {
   if n < 0 { n = 0 }
   if n > len(c.buttons) { n = len(c.buttons) - 1 }

   c.index = n
}


func (c Choice) ToRenderField(x,y, width int) renderer.RenderField{

   fActive := func(c string) renderer.Renderable {
      return renderer.GenerateNode(c, "active")
   }
   f := func(c string) renderer.Renderable { return renderer.GenerateNoRenderNode(c)}


   rl := make([]renderer.RenderLine, len(c.buttons) + 2)
   rl[0] = renderer.GenerateLine(
      width - 2,
      []renderer.Renderable{renderer.GenerateNoRenderNode(c.Header)},
   )
   rl[len(c.buttons) + 1] = renderer.GenerateLine(
      width - 2,
      []renderer.Renderable{renderer.GenerateNoRenderNode(c.Footer)},
   ) 

   for i := 0; i < len(c.buttons); i++{
      var c1 renderer.Renderable
      var c2 renderer.Renderable
      if i == c.index {
         c1 = fActive(c.buttons[i])
         c2 = f(" " + c.labels[i])
      } else {
         c1 = f(c.buttons[i])
         c2 = f(" " + c.labels[i])
      }

      rl[i + 1] = renderer.GenerateLine(width-2, []renderer.Renderable{c1, c2})
   }

   return renderer.GenerateField(rl)
}
func (c Choice) GetSelection() string { return c.labels[c.index]}

