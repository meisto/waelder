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

var styleMap map[string]lipgloss.Style = make(map[string]lipgloss.Style)

func SetupStylemap(style string) {
	// TODO: regex for validation
	// TODO: add italic, underscore, bold, ... support
	// TODO: make colorscheme changeable

	entries := strings.Split(style, "\n")
	for _, e := range entries {
      if e == "" || strings.HasPrefix(e, "//"){ continue }

		fields := strings.Split(e, ",")

		if len(fields) < 3 {
			log.Fatal("[ERROR] Code: 14204875, ", len(fields))
		}

		style := lipgloss.NewStyle()

		if fields[1] != "" {
			style = style.Foreground(lipgloss.Color(fields[1]))
		}

		if fields[2] != "" {
			style = style.Background(lipgloss.Color(fields[2]))
		}

      if len(fields) > 3 {
         for i := 3; i < len(fields); i++ {
            switch fields[i] {
               case "bold":         style = style.Bold(true)
               case "faint":        style = style.Faint(true)
               case "italic":       style = style.Italic(true)
               case "crossout":     style = style.Strikethrough(true)
               case "underline":    style = style.Underline(true)
            }
         }
      }


		styleMap[fields[0]] = style
	}
}

func GetStyle(key string) lipgloss.Style {
	value, isThere := styleMap[key]

	if isThere {
		return value
	} else {
		r, _ := styleMap["unknown"]
		return r
	}
}

func PrintAvailableStyles() {
	for i := range styleMap {
		println(
			styleMap[i].Render(i),
			styleMap[i].GetForeground(),
			styleMap[i].GetBackground(),
		)
	}
}

const (
	cBackground string = "#333333"
)
