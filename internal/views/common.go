// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Sun 15 Jan 2023 10:39:05 PM CET
// Description: -
// ======================================================================
package views

import (
   "math"
   "fmt"
   "strings"

   "dntui/internal/config"
)


func FormatHealthString(
   width int,
   percentage float64,
   isActive bool,
) string {
   // fmt.Sprintf("%s", bar.ViewAs(hpPercentage)),
   // "♡♥❤ ",
   heartIcon := "♥" // "♡"

   fullHearts := int(math.Round(percentage * float64(width)))
   emptyHearts := width - fullHearts

   var a string
   var b string
   if !isActive {
      a = config.StyleHealthBarFull.Render(strings.Repeat(heartIcon, fullHearts))
      b = config.StyleHealthBarEmpty.Render(strings.Repeat(heartIcon, emptyHearts))
   } else {
      a = config.StyleHealthBarFullA.Render(strings.Repeat(heartIcon, fullHearts))
      b = config.StyleHealthBarEmptyA.Render(strings.Repeat(heartIcon, emptyHearts))
   }

   return a + b 
}

func FormatString(
   content string,
   isActive bool,
   isDead bool,
) string {
   if isActive {
      content = config.StyleSelected.Render(content)
   } else if isDead {
      content = config.StyleDead.Render(content)
   } else {
      content = config.StyleDefault.Render(content)
   }

   return content
}

func PopUp (
   content []string,
   buttons []string,
   width int,
   height int,
) string {

   // ┌└┘┐─│

   s := "┌" + strings.Repeat("─", width - 2) + "┐\n"

   for _, c := range(content) {
      s += "│ " + c + " │\n"
   }


   s += "└" + strings.Repeat("─", width - 2) + "┘"

   return s
}

func SelectionBox(
   content  []string,
   selected []bool,
   width    int,
   offset   string,
) []string {
   var s []string

   for i, x := range(content) {

      checked := " "
      if selected[i] { 
         checked = "X"
         checked = config.StyleGreen.Render(checked)
      }

      s = append(s, fmt.Sprintf("%s(%2d) %s", offset, i, checked) + 
         " " + x + 
         strings.Repeat(" ", width - 7 - len(x) - len(offset)))
   }

   return s
}
