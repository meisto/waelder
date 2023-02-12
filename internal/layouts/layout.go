// ======================================================================
// Author: meisto
// Creation Date: Wed 18 Jan 2023 03:22:02 AM CET
// Description: -
// ======================================================================
package layouts

import (
   "strings"

	"github.com/muesli/termenv"

	ds "waelder/internal/datastructures"
	"waelder/internal/modes"
)

type Layout struct {
	TotalHeight int
	TotalWidth  int
	Fields      []Field
   output      *termenv.Output
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

