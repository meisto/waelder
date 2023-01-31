// ======================================================================
// Author: meisto
// Creation Date: Mon 21 Nov 2022 03:13:28 PM CET
// Description: -
// ======================================================================
package datastructures

import (
   "waelder/internal/renderer"
)

type Affiliation int64

const (
	Player  Affiliation = iota
	Ally    Affiliation = iota
	Enemy   Affiliation = iota
	Neutral Affiliation = iota
)

type CharacterStats struct {
	Hp          int
	Max_hp      int
	Initiative  int
	Armour      int
   HasReaction bool
}

type Character struct {
	Name        string
	Affiliation Affiliation
	Race        string
	Subrace     string
	Class       string

	Stats CharacterStats
}

func (ch Character) GetStyledName() renderer.Renderable {
   if ch.Affiliation == Ally {
      return renderer.GenerateNode(ch.Name, "alliedUnit")
   }
   if ch.Affiliation == Neutral {
      return renderer.GenerateNode(ch.Name, "neutralUnit")
   }
   if ch.Affiliation == Enemy {
      return renderer.GenerateNode(ch.Name, "enemyUnit")
   }
   if ch.Affiliation == Player {
      return renderer.GenerateNode(ch.Name, "playerUnit")
   }

   return renderer.GenerateNoRenderNode(ch.Name)
}
