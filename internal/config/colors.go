// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Thu 17 Nov 2022 10:45:49 PM CET
// Description: -
// ======================================================================
package config

import (
   "github.com/charmbracelet/lipgloss"
)


const (
   Healthcolor100 string   = "#00FF00"
   Healthcolor75 string    = "#F3F034"
   Healthcolor50 string    = "#DE7007"
   Healthcolor25 string    = "#FF0000"

   cBackground string   = "#333333"
)

var (
   Style1 = lipgloss.NewStyle().
      Background(lipgloss.Color("#84AF87")).
      Foreground(lipgloss.Color("#000000"))

   StyleDefault = lipgloss.NewStyle().
      Foreground(lipgloss.Color("#DDDDDD"))

   StyleDead = lipgloss.NewStyle().
      Foreground(lipgloss.Color("#444444"))
   StyleSelected = lipgloss.NewStyle().
      Foreground(lipgloss.Color("#333333")).
      Background(lipgloss.Color("#CCCCCC"))


   StyleHealthBarFull = lipgloss.NewStyle().
      Foreground(lipgloss.Color("#FF0000"))
   StyleHealthBarEmpty = lipgloss.NewStyle().
      Foreground(lipgloss.Color("#999999"))
   // Active variants
   StyleHealthBarFullA = lipgloss.NewStyle().
      Foreground(lipgloss.Color("#FF0000")).
      Background(lipgloss.Color("#CCCCCC"))
   StyleHealthBarEmptyA = lipgloss.NewStyle().
      Foreground(lipgloss.Color("#333333")).
      Background(lipgloss.Color("#CCCCCC"))

   StyleGreen = lipgloss.NewStyle().Foreground(lipgloss.Color("#009900"))
   StyleDarkRedBg = lipgloss.NewStyle().Background(lipgloss.Color("#990000"))
)

