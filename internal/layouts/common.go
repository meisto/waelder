// ======================================================================
// Author: meisto
// Creation Date: Sun 15 Jan 2023 10:39:05 PM CET
// Description: -
// ======================================================================
package layouts
/*

import (
   "strings"

	"github.com/muesli/termenv"

	"waelder/internal/renderer"
   io "waelder/internal/wio"
   ds "waelder/internal/datastructures"
   "waelder/internal/modes"
)

func PopUp(
	layout Layout,
	output *termenv.Output,
	content renderer.RenderField,
	x int,
	y int,
   data ds.Data,
) {

   height := content.GetHeight()
   width := content.GetWidth()

   y = y - height / 2 - 1
   
	f := Field{
		x:           x,
		y:           y,
		width:       width + 2,
		height:      height + 2,
		borders:     [4]bool{true, true, true, true},
		borderStyle:   DoubleBorderStyle,
      padding:       [4]int{0,0,0,0},
      content:       content,
      mode:          modes.NoMode,
	}

   // Clear content of popup
   for i := y; i < y + height; i++ {
      output.MoveCursor(i + 2,x + 2)
      output.WriteString(strings.Repeat(" ", width))
   }

	f.DrawBorder(output)
   f.DrawContent(output, data)
}

func ReadLinePopUp(
	layout Layout,
	content renderer.RenderField,
	x int,
	y int,
   data ds.Data,
) string {
   height := content.GetHeight()
   width := content.GetWidth()

	f := Field{
		x:             x,
		y:             y,
		width:         width + 2,
		height:        height + 2,
		borders:       [4]bool{true, true, true, true},
		borderStyle:   DoubleBorderStyle,
      padding:       [4]int{0,0,0,0},
      content:       content,
      mode:          modes.NoMode,
	}


   // Clear content of popup
   for i := y + 1; i < y + height - 1; i++ {
      output.MoveCursor(i + 1,x + 2)
      output.WriteString(strings.Repeat(" ", width - 2))
   }

	f.DrawBorder(output)
   f.DrawContent(output, data)
   
   return <- io.ReadLine(true)
}
*/
