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
			{
				x:           0,
				y:           0,
				height:      height,
				width:       width,
				mode:        modes.ActiveMode,
				borders:     [4]bool{true, true, true, true},
				borderStyle: FancyBorderStyle,
            scrollIndex: -1,
			},
		},
	}

}

func TwoOneHorizontalSplit(height, width int) Layout {
	width -= 1

	l1 := int(float64(width) * 0.5)
	l2 := int(float64(height) * 0.6)

	// ─

	return Layout{
		TotalHeight: height,
		TotalWidth:  width,
		Fields: []Field{
			{
				x:             0,
				y:             0,
				height:        l2,
				width:         l1,
				mode:          modes.ActiveMode,
				borders:       [4]bool{true, false, true, false},
            padding:       [4]int{0, 0, 0, 0},
				borderStyle: DoubleBorderStyle,
            scrollIndex:   -1,
			},
			{
				x:           l1,
				y:           0,
				height:      l2,
				width:       width - l1 + 1,
				mode:          modes.MdViewMode,
				borders:       [4]bool{true, true, true, true},
            padding:       [4]int{1, 2, 1, 2},
				borderStyle:   DoubleBorderStyle,
            scrollIndex:   0,
			},
			{
				x:           0,
				y:           l2 + 0,
				height:      height - l2,
				width:       width,
				mode:        modes.ActionMode,
				borders:     [4]bool{false, false, true, false},
            padding:       [4]int{0, 0, 0, 0},
				borderStyle: DoubleBorderStyle,
            scrollIndex: -1,
			},
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
