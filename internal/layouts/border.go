// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Fri 27 Jan 2023 12:16:44 AM CET
// Description: Some different border styles
// ======================================================================
package layouts

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
