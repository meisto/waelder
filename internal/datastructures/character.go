// ======================================================================
// Author: meisto
// Creation Date: Mon 21 Nov 2022 03:13:28 PM CET
// Description: -
// ======================================================================
package datastructures

import (
// tea "github.com/charmbracelet/bubbletea"
)

type Affiliation int64

const (
	Player  Affiliation = iota
	Ally    Affiliation = iota
	Enemy   Affiliation = iota
	Neutral Affiliation = iota
)

type CharacterStats struct {
	Hp         int
	Max_hp     int
	Initiative int
	Armout     int
}

type Character struct {
	Name        string
	Affiliation Affiliation
	Race        string
	Subrace     string
	Class       string

	Stats CharacterStats
}
