// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Wed 01 Feb 2023 02:59:56 AM CET
// Description: -
// ======================================================================
package datastructures

// Struct to hold a combat log
type CombatLog struct {
	PreviousRounds []Round
	Current        Round
}

type Round struct {
	RoundNumber     int
	Done            []string
	ActiveCharacter string
	Pending         []string
   Actions         []Action
}

func (round Round) IsDone() bool { return len(round.Pending) == 0 }
func (round Round) NumberCombatants() int {
   return len(round.Done) + 1 + len(round.Pending)
}
