// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Sun 15 Jan 2023 10:39:05 PM CET
// Description: -
// ======================================================================
package views

import (
   "log"
   "math"
   "fmt"
   "strings"

   "github.com/charmbracelet/lipgloss"

   "waelder/internal/config"
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

type ColorHelper struct {
   Index int
   Style lipgloss.Style
}

type PopupBorder struct {
   topLeftCorner     string
   topRightCorner    string
   lowerRightCorner  string
   lowerLeftCorner   string
   topBorder         string
   rightBorder       string
   bottomBorder      string
   leftBorder        string
}

var DefaultBorder PopupBorder = PopupBorder {
   topLeftCorner: "┌",
   topRightCorner: "┐",
   lowerRightCorner: "┘",
   lowerLeftCorner: "└",
   topBorder: "─",
   rightBorder: "│",
   bottomBorder: "─",
   leftBorder: "│",
}

func StylePopupBorder(
   b     PopupBorder,
   style lipgloss.Style,
) PopupBorder {
   return PopupBorder {
      style.Render(b.topLeftCorner),
      style.Render(b.topRightCorner),
      style.Render(b.lowerRightCorner),
      style.Render(b.lowerLeftCorner),
      style.Render(b.topBorder),
      style.Render(b.rightBorder),
      style.Render(b.bottomBorder),
      style.Render(b.leftBorder),
   }
}

func PopUp (
   content [][]string,
   ch [][]ColorHelper,
   bs PopupBorder,
   offsetTop int,
   offsetLeft int,
) string {

   // Calculate width of the popup window.
   width    := 0
   for _, line := range(content) {
      l := -1
      for _, i := range(line) { l += len(i) + 1}
      if (l + 4) > width {width = l + 4}
   }

   // Assemble main string
   strMain  := ""
   for i, line := range(content) {
      
      // Get length of line without padding
      l := -1 // For spaces between elements
      for _, i := range(line) { l += len(i) + 1}
      if l == -1 {l = 0}   // For empty lines
      if (l + 4) > width {width = l + 4}

      // Apply styles
      if len(content) != len(ch) {log.Fatal("[ERROR] Code:12302913")}
      styles := ch[i]
      for _, x := range(styles) {line[x.Index] = x.Style.Render(line[x.Index])}
      
      // Add padding
      paddingC := 0
      if l < width - 4 {
         paddingC = width - l - 4
      }
   

      strMain += 
         strings.Repeat(" ", offsetLeft) +   // Offset
         bs.leftBorder + " " +               // Left inner padding
         strings.Join(line, " ") +           // Join styled elements
         strings.Repeat(" ", paddingC) +     // Add Padding
         " " +                               // Right inner padding
         bs.rightBorder + "\n"               // Right border
   }

   // Assemble string
   return strings.Repeat("\n", offsetTop) +     // Top Bar
      strings.Repeat(" ", offsetLeft) + 
      bs.topLeftCorner + 
      strings.Repeat(bs.topBorder, width - 2) +
      bs.topRightCorner +
      "\n" + 
      strMain +                                 // Content
      strings.Repeat(" ", offsetLeft) +         // Bottom Bar
      bs.lowerLeftCorner +
      strings.Repeat(bs.bottomBorder, width - 2) +
      bs.lowerRightCorner
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
