// ======================================================================
// Author: meisto
// Creation Date: Fri 27 Jan 2023 12:26:04 AM CET
// Description: -
// ======================================================================
package layouts

import (
	"fmt"
	"strings"

	// tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"

	ds "waelder/internal/datastructures"
	"waelder/internal/modes"
	"waelder/internal/config"
)

type ScrollField struct {
	x           int
	y           int
	width       int
	height      int

	borders     [3]bool
   padding     [4]int
	borderStyle BorderStyle
}

func ScrollUp() {
      index = index - 1

   if index < 0 { index = 0 }
}
func ScrollDown() {
      index += 1
}
var index int = 0

func (f ScrollField) GetMode() modes.Mode { return modes.MdViewMode }
func (f *ScrollField) SetBorder(f2 *ScrollField, bs BorderStyle) { f2.borderStyle = bs }
func (f ScrollField) GetBorder() BorderStyle { return f.borderStyle }
func GetIndex() int { return index }
func (f ScrollField) DrawBorder(output *termenv.Output) {
   scrollPercentage := float64(index) / float64(modes.GetMarkdownLength() )

	// Get border style elements
	t := f.borderStyle.upperBorder
	r := f.borderStyle.rightBorder
	b := f.borderStyle.lowerBorder
	l := f.borderStyle.leftBorder

	tl := f.borderStyle.ulCorner
	tr := f.borderStyle.urCorner
	br := f.borderStyle.lrCorner
	bl := f.borderStyle.llCorner

	g := func(x, y int, style string) {
		output.MoveCursor(x+1, y+1)
		fmt.Print(output.String(style))
	}

	// Borders
	lengthL := f.width - 1
	if f.borders[2] { lengthL -= 1 }

	lengthR := f.width - 1
	if f.borders[2] { lengthR -= 1 }

   // Draw top border
	if f.borders[0] { g(f.y, f.x+1, strings.Repeat(t, lengthL)) }
   
   // Draw right border
		for i := 1; i < f.height-1; i++ {
         
         s := r
         if modes.GetMarkdownLength() > f.height - 2 && 
          ((i <= 1 && scrollPercentage <= 0.1) || (i >= f.height - 2 && scrollPercentage >= 0.95) || i == int(float64(f.height - 1) * scrollPercentage)) {
            s = config.GetStyle("darkRedFg").Render("╬")
         }

         g(f.y+i, f.x+f.width-1, s) 
      }

   // Draw bottom border
	if f.borders[1] { g(f.y+f.height-1, f.x+1, strings.Repeat(b, lengthR)) }

   // Draw left border
	if f.borders[2] {
		for i := 1; i < f.height-1; i++ { g(f.y+i, f.x, l) }
	}

	// Corners
	if f.borders[0] {
      g(f.y, f.x+f.width-1, tr)
		if !f.borders[2] {
			g(f.y, f.x, t)
		}
	}
   if f.borders[1] {
      g(f.y+f.height-1, f.x+f.width-1, br)
   } else {
      g(f.y+f.height-1, f.x+f.width-1, r)
   }
   if !f.borders[0] {
      g(f.y, f.x+f.width-1, r)
   }
	if f.borders[1] {
		if f.borders[2] {
			g(f.y+f.height-1, f.x, bl)
		} else {
			g(f.y+f.height-1, f.x, b)
		}
	}
	if f.borders[2] {
		if f.borders[0] {
			g(f.y, f.x, tl)
		} else {
			g(f.y, f.x, l)
		}
		if !f.borders[1] {
			g(f.y+f.height-1, f.x, l)
		}
	}

}
func (f ScrollField) DrawContent(output *termenv.Output, data ds.Data) {
	// Remainig width
	h := f.height  - f.padding[0] - f.padding[2]
	w := f.width   - f.padding[1] - f.padding[3] - 1

	// Offsets
	vOff := f.padding[0]
	hOff := f.padding[3]

   // Factor in borders
	if f.borders[0] { vOff += 1; h -= 1 }
	if f.borders[1] { h -= 1 }
	if f.borders[2] { hOff += 1; w -= 1 }
   
	modes.ModeLookup[modes.MdViewMode](output, data, f.x+hOff, f.y+vOff, h, w)
}
