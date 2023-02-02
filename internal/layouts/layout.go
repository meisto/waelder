// ======================================================================
// Author: meisto
// Creation Date: Wed 18 Jan 2023 03:22:02 AM CET
// Description: -
// ======================================================================
package layouts

import (
	// tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"

	ds "waelder/internal/datastructures"
	"waelder/internal/modes"
)

type Layout struct {
	TotalHeight int
	TotalWidth  int
	Fields      []Field
}

func FullScreen(height, width int) Layout {
	return Layout{
		TotalHeight: height,
		TotalWidth:  width,
		Fields: []Field{
			GetField( 0, 0, height, width, modes.ActiveMode, [4]int{0,0,0,0}, -1, [4]bool{true, true, true, true}, FancyBorderStyle),
		},
	}
}


   /*
funcGetField(x, y, width, height int, mode modes.Mode,
   padding [4]int, scrollIndex int, borders [4]bool, borderStyle BorderStyle) Field {
}*/

func TwoOneHorizontalSplit(height, width int) Layout {
	width -= 1

	l1 := int(float64(width) * 0.5)
	l2 := int(float64(height) * 0.8)

	// ─

	return Layout{
		TotalHeight: height,
		TotalWidth:  width,
		Fields: []Field{
			GetField(
				0,
			   0,
				l1,
				l2,
				modes.ActiveMode,
            [4]int{0,0,0,0},
            100,
				[4]bool{true, false, true, false},
				DoubleBorderStyle,
			),
			GetField(
			   l1,
				0,
				width - l1 + 1,
				l2,
				modes.MdViewMode,
            [4]int{1,2,1,2},
            0,
				[4]bool{true, true, true, true},
				DoubleBorderStyle,
			),
			GetField(
				0,
            l2,
				width,
				height - l2 - 1,
				modes.ActionMode,
            [4]int{0,0,0,0},
            -1,
				[4]bool{false, false, true, false},
				DoubleBorderStyle,
			),
			GetField(
				0,
            height - 1,
				width,
				1,
				modes.HelpMode,
            [4]int{0,0,0,0},
            -1,
				[4]bool{false, false, false, false},
				DoubleBorderStyle,
			),
		},
	}
}

func (lay Layout) DrawBorders(output *termenv.Output) {
	for _, f := range lay.Fields {
		f.DrawBorder(output)
	}
}

func (lay *Layout) Display(output *termenv.Output, data ds.Data) {
   for i := 0; i < len(lay.Fields); i++ { 
      lay.Fields[i].DrawContent(output, data) 
   }
}

func (lay *Layout) UpdateMode(output *termenv.Output, data ds.Data, mode modes.Mode) {
	for _, f := range lay.Fields {
		if f.GetMode() == mode { f.DrawContent(output, data) }
	}
}

func (lay *Layout) Reset(output *termenv.Output, data ds.Data) {
   output.ClearScreen()
	lay.DrawBorders(output)
	lay.Display(output, data)
	output.MoveCursor(lay.TotalHeight+1, 0)
}
