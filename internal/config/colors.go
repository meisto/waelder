// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Thu 17 Nov 2022 10:45:49 PM CET
// Description: -
// ======================================================================
package config

import (
   "log"
   "strings"

   "github.com/charmbracelet/lipgloss"
)

const defaultScheme string = 
   "style1,#000000,#84AF87\n" + 
   "default,#DDDDDD,\n" +
   "dead,#444444,\n" +
   "selected,#333333,#CCCCCC\n" +
   "healthBarFull,#FF0000,\n" + 
   "healthBarEmpty,#999999,\n" + 
   "healthBarFullActive,#FF0000,#CCCCCC\n" + 
   "healthBarEmptyActive,#333333,#CCCCCC\n" + 
   "green,#009900,\n" +
   "darkRedBg,,#990000"


var styleMap map[string]lipgloss.Style = make(map[string]lipgloss.Style)
func SetupStylemap() {
   // TODO: regex for validation
   // TODO: add italic, underscore, bold, ... support
   // TODO: make colorscheme changeable

   entries := strings.Split(defaultScheme, "\n")
   for _, e := range(entries) {
      fields := strings.Split(e, ",")

      if len(fields) < 3 { log.Fatal("[ERROR] Code: 14204875") } 

      style := lipgloss.NewStyle()

      if fields[1] != "" {
         style = style.Foreground(lipgloss.Color(fields[1]))
      }

      if fields[2] != "" {
         style = style.Background(lipgloss.Color(fields[2]))
      }

      styleMap[fields[0]] = style
   }
}

func GetStyle(key string) lipgloss.Style {
   value, isThere := styleMap[key]

   if isThere {
      return value
   } else {
      r, _ := styleMap["default"]
      return r
   }
}

func PrintAvailableStyles() {
   for i := range(styleMap) {
      println(
         styleMap[i].Render(i), 
         styleMap[i].GetForeground(),
         styleMap[i].GetBackground(),
      )
   }
}

const (
   cBackground string   = "#333333"
)


