// ======================================================================
// Author: meisto
// Creation Date: Sun 15 Jan 2023 10:39:05 PM CET
// Description: -
// ======================================================================
package modes

import (
	"math"
	"strings"

	"waelder/internal/renderer"
)

func FormatHealthString(
	width int,
	percentage float64,
	isActive bool,
	heartIcon string,
) renderer.RenderLine {

	fullHearts := int(math.Round(percentage * float64(width)))
	emptyHearts := width - fullHearts

	var a renderer.Renderable
	var b renderer.Renderable
	if !isActive {
		a = renderer.GenerateNode(
			strings.Repeat(heartIcon, fullHearts),
			"healthBarFull",
		)
		b = renderer.GenerateNode(
			strings.Repeat(heartIcon, emptyHearts),
			"healthBarEmpty",
		)
	} else {
		a = renderer.GenerateNode(
			strings.Repeat(heartIcon, fullHearts),
			"healthBarFullActive",
		)
		b = renderer.GenerateNode(
			strings.Repeat(heartIcon, emptyHearts),
			"healthBarEmptyActive",
		)
	}

	return renderer.GenerateLine(width, []renderer.Renderable{a, b})
}
