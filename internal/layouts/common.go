// ======================================================================
// Author: meisto
// Creation Date: Sun 15 Jan 2023 10:39:05 PM CET
// Description: -
// ======================================================================
package layouts

import (
   "strings"

	"github.com/muesli/termenv"

	"waelder/internal/renderer"
   io "waelder/internal/wio"
)

func PopUp(
	output *termenv.Output,
	content renderer.RenderField,
	x int,
	y int,
	width int,
	height int,
) {
	f := Field{
		x:           x,
		y:           y,
		width:       width,
		height:      height,
		borders:     [4]bool{true, true, true, true},
		borderStyle: DoubleBorderStyle,
	}

	f.DrawBorder(output)

   // Clear content of popup
   for i := y + 1; i < y + height - 1; i++ {
      output.MoveCursor(i + 1,x + 2)
      output.WriteString(strings.Repeat(" ", width - 2))
   }

	content.RenderBlock(output, x+1, y+1, height-2, true, 10000)
}

func ReadLinePopUp(
	output *termenv.Output,
	content renderer.RenderField,
	x int,
	y int,
	width int,
	height int,
) string {
	f := Field{
		x:           x,
		y:           y,
		width:       width,
		height:      height,
		borders:     [4]bool{true, true, true, true},
		borderStyle: DoubleBorderStyle,
	}

	f.DrawBorder(output)

   // Clear content of popup
   for i := y + 1; i < y + height - 1; i++ {
      output.MoveCursor(i + 1,x + 2)
      output.WriteString(strings.Repeat(" ", width - 2))
   }

	content.RenderBlock(output, x, y, height-2, true, 10000)
   
   return <- io.ReadLine(true)
}



