// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Sun 15 Jan 2023 11:45:05 PM CET
// Description: -
// ======================================================================
package datastructures

import (
   "fmt"
   "log"
	"sort"
)

// Root data struct
type Data struct {
	Players  []string
	Allies   []string
	Enemies  []string
	Neutrals []string

	CombatLog CombatLog
   characters  map[string]Character
}
func GetData() Data { 
   return Data {
      []string{},
      []string{},
      []string{},
      []string{},

      CombatLog{
         PreviousRounds:   []Round{},
         Current: Round{
            RoundNumber:   -1,
            Done:             []string{},
            ActiveCharacter:  "",
            Pending:          []string{},
            Actions:          []Action{},
         },
      },
      make(map[string]Character),
   } 
}
func (d *Data) AddPlayer(c Character) {
   d.Players = append(d.Players, c.Name)
   d.characters[c.Name] = c
}
func (d *Data) AddAlly(c Character) {
   d.Allies = append(d.Allies, c.Name)
   d.characters[c.Name] = c
}
func (d *Data) AddEnemy(c Character) {
   d.Enemies = append(d.Enemies, c.Name)
   d.characters[c.Name] = c
}
func (d *Data) AddNeutral(c Character) {
   d.Neutrals = append(d.Neutrals, c.Name)
   d.characters[c.Name] = c
}


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

/*
*

	Progress the round.

	Returns true if a state transition was done, false if round is done

*
*/
func (data *Data) Step() {

   round := &data.CombatLog.Current

	// Break
	if round.IsDone() {
      // Round is done, prepare next round and return
		data.PrepareNextRound()
      return
	}

	round.Done = append(round.Done, round.ActiveCharacter)

	// Sort by initiative
	comp := func(i, j int) bool {
		return data.characters[round.Pending[i]].Stats.Initiative > 
         data.characters[round.Pending[i]].Stats.Initiative
	}
	sort.Slice(round.Pending, comp)

	round.ActiveCharacter = round.Pending[0]
	round.Pending = round.Pending[1:]
}

func (data *Data) PrepareNextRound() {
   if !data.CombatLog.Current.IsDone() { return }


   currentRound   := data.CombatLog.Current

   // Only add current round to log if it was actually initialized
   if data.CombatLog.Current.RoundNumber >= 0 {
      data.CombatLog.PreviousRounds = 
         append(data.CombatLog.PreviousRounds, currentRound)
   }
   // Copy slice as not to alter the old order
   pending := data.Players
   pending = append(pending, data.Allies...)
   pending = append(pending, data.Enemies...)
   pending = append(pending, data.Neutrals...)


	// Sort by initiative
	sort.Slice(pending, func(i, j int) bool {
		return data.characters[pending[i]].Stats.Initiative > 
         data.characters[pending[i]].Stats.Initiative
	})

   log.Print(fmt.Sprintf("Prepared round number %d", currentRound.RoundNumber + 1))

	data.CombatLog.Current = Round{
		RoundNumber:     currentRound.RoundNumber + 1,
		Done:             []string{},
		ActiveCharacter:  pending[0],
		Pending:          pending[1:],
      Actions:          []Action{},
	}
}

func (data Data) GetCharacter(name string) Character {
   return data.characters[name]
}
func (data *Data) AddAction(action Action) {
   data.CombatLog.Current.Actions = 
      append(data.CombatLog.Current.Actions, action)

   for _, i := range(action.GetTargets()) {
      data.characters[i] = action.Apply(data.characters[i])
   }
}
