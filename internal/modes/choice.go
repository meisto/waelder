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
   buttons  []string
   labels   []renderer.Renderable
}
func GetChoice(header string, buttons []string, labels []renderer.Renderable) Choice {
   return Choice{0, header, buttons, labels}
}
func (c *Choice) Forward() {
   c.index += 1
   if c.index >= len(c.buttons) { c.index = 0}
}
func (c *Choice) Backward() {
   c.index -= 1
   if c.index < 0 { c.index = len(c.buttons) - 1}
}

func (c *Choice) GoTo(n int) {
   if n < 0 { n = 0 }
   if n >= len(c.buttons) { n = len(c.buttons) - 1 }

   c.index = n
}


func (c Choice) ToRenderField(x,y, width int) renderer.RenderField{

   fActive := func(c string) renderer.Renderable {
      return renderer.GenerateNode(c, "active")
   }


   rl := make([]renderer.RenderLine, len(c.buttons) + 1)
   rl[0] = renderer.GenerateLine(
      width - 2,
      []renderer.Renderable{renderer.GenerateNoRenderNode(c.Header)},
   )

   for i := 0; i < len(c.buttons); i++{
      var c1 renderer.Renderable
      var c2 renderer.Renderable
      if i == c.index {
         c1 = fActive(c.buttons[i])
         c2 = c.labels[i]
      } else {
         c1 = renderer.GenerateNoRenderNode(c.buttons[i])
         c2 = c.labels[i]
      }

      rl[i + 1] = renderer.GenerateLine(width-2, []renderer.Renderable{c1, c2})
   }

   return renderer.GenerateField(rl)
}

func (c Choice) GetIndex() int { return c.index }
