// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Wed 18 Jan 2023 03:22:02 AM CET
// Description: -
// ======================================================================
package layouts

import (
   // tea "github.com/charmbracelet/bubbletea"
   "github.com/muesli/termenv"

   "waelder/internal/modes"
   ds "waelder/internal/datastructures"


)


type Layout struct {
   TotalHeight int
   TotalWidth  int
   fields      []Field
}


func FullScreen(height, width int) Layout {
   return Layout {
      TotalHeight:   height,
      TotalWidth:    width,
      fields: []Field {
         {
            x:       0,
            y:       0,
            height:  height,
            width:   width,
            mode: modes.ActiveMode,
            borders: [4]bool{true, true, true, true},
            borderStyle: FancyBorderStyle,
         },
      },
   }
   
}

func TwoOneHorizontalSplit(height, width int) Layout {
   width -= 1

   l1 := int(float64(width) * 0.5)
   l2 := int(float64(height) * 0.6)

   // ─

   return Layout {
      TotalHeight:   height,
      TotalWidth:    width,
      fields: []Field {
         {
            x:       0,
            y:       0,
            height:  l2,
            width:   l1,
            mode: modes.ActiveMode,
            borders: [4]bool{true, false, true, false},
            borderStyle: FancyBorderStyle,
         },
         {
            x:       l1,
            y:       0,
            height:  l2,
            width:   width - l1 + 1,
            mode: modes.ActiveMode,
            borders: [4]bool{true, false, true, true},
            borderStyle: FancyBorderStyle,
         },
         {
            x:       0,
            y:       l2 + 0,
            height:  height - l2,
            width:   width,
            mode: modes.ActionMode,
            borders: [4]bool{false, false, true, false},
            borderStyle: FancyBorderStyle,
         },
      },
   }
}

func (lay Layout) DrawBorders(output *termenv.Output) {
   for _, f := range(lay.fields) { f.DrawBorder(output) }
}


func (lay *Layout) Display(output *termenv.Output, data ds.Data) {
   for _, f := range(lay.fields) { f.DrawContent(output, data) }
   output.MoveCursor(lay.TotalHeight - 1, 1)
}

func (lay *Layout) UpdateMode(output *termenv.Output, data ds.Data, mode modes.Mode) {
   for _, f := range(lay.fields) { 
      if f.mode == mode { 
         f.DrawContent(output, data)
      }
   }
}
