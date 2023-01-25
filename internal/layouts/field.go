// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Wed 18 Jan 2023 03:22:02 AM CET
// Description: -
// ======================================================================
package layouts

import (
   "fmt"
   "strings"

   // tea "github.com/charmbracelet/bubbletea"
   "github.com/muesli/termenv"

   "waelder/internal/modes"
   ds "waelder/internal/datastructures"
)

type Field struct {
   x           int
   y           int
   width       int
   height      int
   mode        modes.Mode
   borders     [4]bool
   borderStyle BorderStyle
}

type BorderStyle struct {
   ulCorner string
   urCorner string
   llCorner string
   lrCorner string
   upperBorder string
   leftBorder string
   lowerBorder string
   rightBorder string
}

func (f Field) DrawBorder(output *termenv.Output) {

   // Get border style elements
   t  := f.borderStyle.upperBorder
   r  := f.borderStyle.rightBorder
   b  := f.borderStyle.lowerBorder
   l  := f.borderStyle.leftBorder

   tl := f.borderStyle.ulCorner
   tr := f.borderStyle.urCorner
   br := f.borderStyle.lrCorner
   bl := f.borderStyle.llCorner

   g := func(x,y int, style string) {
      output.MoveCursor(x + 1, y + 1)
      fmt.Print(output.String(style))
   }
   
   // Borders
   lengthL := f.width
   if f.borders[1] { lengthL -= 1 }
   if f.borders[3] { lengthL -= 1 }
   lengthR := f.width
   if f.borders[1] { lengthR -= 1 }
   if f.borders[3] { lengthR -= 1 }

   if f.borders[0] { g(f.y, f.x + 1, strings.Repeat(t, lengthL)) }
   if f.borders[1] {
      for i := 1; i < f.height - 1; i++ {
         g(f.y + i, f.x + f.width - 1, r)
      }
   }
   if f.borders[2] { g(f.y + f.height - 1, f.x + 1, strings.Repeat(b, lengthR)) }
   if f.borders[3] {
      for i := 1; i < f.height - 1; i++ {
         g(f.y + i, f.x, l)
      }
   }

   // Corners
   if f.borders[0] {
      if f.borders[1] { g(f.y, f.x + f.width - 1, tr) } else { g(f.y, f.x + f.width - 1, t) }
      if !f.borders[3] { g(f.y, f.x, t) }
   }
   if f.borders[1] {
      if f.borders[2] { g(f.y + f.height - 1, f.x + f.width - 1, br) } else { g(f.y + f.height - 1, f.x + f.width - 1, r) }
      if !f.borders[0] { g(f.y, f.x + f.width -1, r) }
   }
   if f.borders[2] {
      if f.borders[3] { g(f.y + f.height - 1, f.x, bl) } else { g(f.y + f.height - 1, f.x, b) }
      if !f.borders[1] { g(f.y + f.height - 1, f.x + f.width -1, b) }
   }
   if f.borders[3] {
      if f.borders[0] { g(f.y, f.x, tl) } else { g(f.y, f.x, l) }
      if !f.borders[2] { g(f.y + f.height - 1, f.x, l) }
   }
}

func (f Field) DrawContent(output *termenv.Output, d ds.Data) {
   mh := modes.ModeLookup[f.mode]

   // Remainig width
   h := f.height
   w := f.width

   // Offsets
   hOff := 0
   vOff := 0

   if f.borders[0] { 
      vOff += 1
      h -= 1
   }
   if f.borders[3] {
      hOff += 1
      w -= 1
   }
   if f.borders[1] {w -= 1}
   if f.borders[2] {h -= 1}

   mh.View(output, d, f.x + hOff, f.y + vOff, h, w)

}


var DefaultBorderStyle BorderStyle = BorderStyle {
   ulCorner:      "+",
   urCorner:		"+",
   llCorner:		"+",
   lrCorner:		"+",
   upperBorder:	"-",
   leftBorder:		"|",
   lowerBorder:   "-",
   rightBorder:   "|",
}

var FancyBorderStyle BorderStyle = BorderStyle {
   ulCorner:      "┌",
   urCorner:		"┐",
   llCorner:		"└",
   lrCorner:		"┘",
   upperBorder:	"─",
   leftBorder:		"│",
   lowerBorder:   "─",
   rightBorder:   "│",
}
