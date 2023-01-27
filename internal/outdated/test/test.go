// ======================================================================
// Author: meisto
// Creation Date: Wed 18 Jan 2023 11:04:34 PM CET
// Description: -
// ======================================================================
package main

import (
	"waelder/internal/config"
	"waelder/internal/renderer"
)

func mainLocal() {
	config.SetupStylemap()
	a := renderer.GenerateNode("123", "style1")

	println(a.Render())
	println()

	println()
	config.PrintAvailableStyles()

	println()
}
