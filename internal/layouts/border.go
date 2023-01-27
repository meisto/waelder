// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Fri 27 Jan 2023 12:16:44 AM CET
// Description: Some different border styles
// ======================================================================
package layouts

import "waelder/internal/config"

type BorderStyle struct {
	ulCorner    string
	urCorner    string
	llCorner    string
	lrCorner    string
	upperBorder string
	leftBorder  string
	lowerBorder string
	rightBorder string
}

func(b BorderStyle) Style(style string) BorderStyle {
   s := config.GetStyle(style)

   return BorderStyle {
      s.Render(b.ulCorner),
      s.Render(b.urCorner),
      s.Render(b.llCorner),
      s.Render(b.lrCorner),
      s.Render(b.upperBorder),
      s.Render(b.leftBorder),
      s.Render(b.lowerBorder),
      s.Render(b.rightBorder),
   }
}

var DefaultBorderStyle BorderStyle = BorderStyle{
	ulCorner:    "+",
	urCorner:    "+",
	llCorner:    "+",
	lrCorner:    "+",
	upperBorder: "-",
	leftBorder:  "|",
	lowerBorder: "-",
	rightBorder: "|",
}

var FancyBorderStyle BorderStyle = BorderStyle{
	ulCorner:    "┌",
	urCorner:    "┐",
	llCorner:    "└",
	lrCorner:    "┘",
	upperBorder: "─",
	leftBorder:  "│",
	lowerBorder: "─",
	rightBorder: "│",
}

var DoubleBorderStyle = BorderStyle{
	ulCorner:    "╔",
	urCorner:    "╗",
	llCorner:    "╚",
	lrCorner:    "╝",
	upperBorder: "═",
	leftBorder:  "║",
	lowerBorder: "═",
	rightBorder: "║",
}
