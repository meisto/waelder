// ======================================================================
// Author: meisto
// Creation Date: Mon 21 Nov 2022 03:13:28 PM CET
// Description: -
// ======================================================================
package datastructures

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

func (ch Character) IsDead() bool { return ch.Stats.Hp <= 0}

func (ch Character) GetStyle() string {
   switch ch.Affiliation {
      case Player: return "playerUnit"
      case Ally: return "allyUnit"
      case Enemy: return "enemyUnit"
      case Neutral: return "neutralUnit"
   }
   return "unknown"
}
