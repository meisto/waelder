// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Tue 17 Jan 2023 05:04:20 PM CET
// Description: -
// ======================================================================
package choicepage

import (
   tea "github.com/charmbracelet/bubbletea"
   "github.com/charmbracelet/lipgloss"

   // "dntui/internal/language"
   "dntui/internal/views"
   ds "dntui/internal/datastructures"
)

type selection []selectionElement
func (s selection) step() selection {if len(s) > 1 {return append(s[1:], s[0])} else {return s}}
func(s selection) apply(d *ds.Data) {s[0].action(d)}

type selectionElement struct {
   display  string
   style    lipgloss.Style
   i1       int
   i2       int
   action   func(d *ds.Data)
}

// View method
func ChoiceView(
   data           ds.Data,
   windowHeight   int,
   windowWidth    int,
) string {

   c := subPageLookup[currentKey]
   
   
   // Generate copies as not to alter the originals
   rows     := make([][]string, len(c.content))
   styles   := make([][]views.ColorHelper, len(c.style))
   for i := 0; i < len(c.content); i++ {
      rows[i] = make([]string, len(c.content[i]))
      copy(rows[i], c.content[i])

      styles[i] = make([]views.ColorHelper, len(c.style[i]))
      copy(styles[i], c.style[i])
   }

   // Apply marker for currently selected element
   if len(c.selection) > 0 {
      t := c.selection[0]
      
      // Change the marker if necessary
      if t.display != "" {rows[t.i1][t.i2] = t.display}

      // Only add style if there is already a style assigned to it
      for i, el := range(styles[t.i1]) {
         if el.Index == t.i2 {
            styles[t.i1][i].Style = t.style
            break
         }
      }
   }
   return views.PopUp(rows, styles, views.DefaultBorder, 0, 2)
}

func ChoiceUpdate(d *ds.Data, msg tea.KeyMsg, rf func()) { 
   returnFunction = rf
   switch msg.String(){
      case "enter":
         subPageLookup[currentKey].selection.apply(d)
      default:
         subPageLookup[currentKey].selection = 
            subPageLookup[currentKey].selection.step()
   }
}

var returnFunction func()
