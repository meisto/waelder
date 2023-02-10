// ======================================================================
// Author: meisto
// Creation Date: Wed 18 Jan 2023 03:22:02 AM CET
// Description: -
// ======================================================================
package layouts

import (
   "strings"

	// tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"

	ds "waelder/internal/datastructures"
	"waelder/internal/modes"
   "waelder/internal/renderer"
)

type Layout struct {
	TotalHeight int
	TotalWidth  int
	Fields      []Field
   output      *termenv.Output
}

func TwoOneHorizontalSplit(height, width int, output *termenv.Output, data ds.Data) Layout {
	width -= 1

	l1 := int(float64(width) * 0.5)
	l2 := int(float64(height) * 0.8)

	// ─
   var content renderer.RenderField

   l := Layout{
		TotalHeight: height,
		TotalWidth:  width,
		Fields: []Field{
			{
            x: 0,
            y: 0,
            width: l1,
            height: l2,
            mode: modes.ActiveMode,
            content: content,
            padding: [4]int{0,0,0,0},
            scrollIndex: 100,
            borders: [4]bool{true,false,true,false},
            borderStyle: DoubleBorderStyle,
            startTop: false,
            output: output,
         },
			{
			   l1,
				0,
				width - l1 + 1,
				l2,
				modes.MdViewMode,
            content,
            [4]int{1,2,1,2},
            0,
				[4]bool{true, true, true, true},
				DoubleBorderStyle,
            false,
            output,
			},
			 {
				0,
            l2,
				width,
				height - l2 - 1,
				modes.ActionMode,
            content,
            [4]int{0,0,0,0},
            -1,
				[4]bool{false, false, true, false},
				DoubleBorderStyle,
            false,
            output,
			},
			{
				0,
            height - 1,
				width,
				1,
				modes.HelpMode,
            content,
            [4]int{0,0,0,0},
            -1,
				[4]bool{false, false, false, false},
				DoubleBorderStyle,
            false,
            output,
         },
		},
      output: output,
	}
   return l 
}

func (lay Layout) DrawBorders() {
	for _, f := range lay.Fields {
		f.DrawBorder()
	}
}

func (lay *Layout) Display(data ds.Data) {
   for i := 0; i < len(lay.Fields); i++ { 
      lay.Fields[i].UpdateContent(data)
      lay.Fields[i].DrawContent() 
   }
}

func (lay *Layout) UpdateMode(data ds.Data, mode modes.Mode) {
	for _, f := range lay.Fields {
		if f.GetMode() == mode { 

         f.UpdateContent(data) 
         f.DrawContent() 
      }
	}
}

func (lay *Layout) Reset(data ds.Data) {
   lay.output.ClearScreen()
	lay.DrawBorders()
	lay.Display(data)
	lay.output.MoveCursor(lay.TotalHeight+1, 0)
}

func (lay *Layout) DisplayPopup(pop PopupField) {

   // Generate actual field from PopopField
	f := Field{
      x: pop.x,
      y: pop.y,
      width: pop.content.GetWidth() + 4,
      height: pop.content.GetHeight() + 2,
      mode: modes.NoMode,
      content: pop.content,
      padding: [4]int{0,1,0,1},
      scrollIndex: 100,
      borders: [4]bool{true,true,true,true},
      borderStyle: DoubleBorderStyle,
      startTop: false,
      output: lay.output,
   }

   // Clear content of popup
   for i := pop.y; i < pop.y + pop.content.GetHeight(); i++ {
      lay.output.MoveCursor(i + 2, pop.x + 2)
      lay.output.WriteString(strings.Repeat(" ", pop.content.GetWidth()))
   }

	f.DrawBorder()
   f.DrawContent()
	lay.output.MoveCursor(lay.TotalHeight+1, 0)
}

