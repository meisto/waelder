// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Tue 17 Jan 2023 05:04:20 PM CET
// Description: -
// ======================================================================
package choicepage

import (
   tea "github.com/charmbracelet/bubbletea"
   "github.com/charmbracelet/lipgloss"

   // "waelder/internal/language"
   "waelder/internal/views"
   ds "waelder/internal/datastructures"
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

   // c := subPageLookup[currentKey]
   

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
