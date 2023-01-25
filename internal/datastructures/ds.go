// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Sun 15 Jan 2023 11:45:05 PM CET
// Description: -
// ======================================================================
package datastructures

import (
   "sort"
)

// Root data struct
type Data struct {
   Players           []Character
   Allies            []Character
   Enemies           []Character
   Neutrals          []Character

   CombatLog         CombatLog
}

// Struct to hold a combat log
type CombatLog struct {
   Actions        []Action
   PreviousRounds []Round
   Current        CurrentRound
}

type Round struct {
   RoundNumber    int
   TurnSequence   []Character
}

type CurrentRound struct {
   RoundNumber       int
   Done              []Character
   ActiveCharacter   Character
   Pending           []Character
}

func (round CurrentRound) IsDone() bool { return len(round.Pending) == 0 }

/**
   Progress the round. 

   Returns true if a state transition was done, false if round is done
**/
func (round *CurrentRound) Step() bool{

   // Break
   if round.IsDone() {return false}
   
   round.Done = append(round.Done, round.ActiveCharacter)

   // Sort by initiative
   comp := func(i,j int) bool {
      return round.Pending[i].Stats.Initiative > round.Pending[j].Stats.Initiative
   }
   sort.Slice(round.Pending, comp)

   round.ActiveCharacter = round.Pending[0]
   round.Pending = round.Pending[1:]

   return true
}

/**
   Return an object representing the (now done) round and a new CurrentRound
   object representing the next round
**/
func (round CurrentRound) GetNextRound() (Round, CurrentRound) {
   a := Round{
      RoundNumber:   round.RoundNumber,
      TurnSequence:  append(round.Done, round.ActiveCharacter),
   }


   // Sort by initiative
   sort.Slice(round.Done, func(i,j int) bool {
      return round.Done[i].Stats.Initiative > round.Done[j].Stats.Initiative
   })


   d := append(round.Done, round.ActiveCharacter)
   b := CurrentRound {
      RoundNumber:      round.RoundNumber + 1,
      ActiveCharacter:  d[0],
      Done:             []Character{},
      Pending:          d[1:],
   }

   return a, b
}

func CreateRound(chars []Character) CurrentRound {
   // Sort by initiative
   sort.Slice(chars, func(i,j int) bool {
      return chars[i].Stats.Initiative > chars[j].Stats.Initiative
   })

   p := []Character{}
   if len(chars) > 1 {p = chars[1:]}
   
   return CurrentRound {
      RoundNumber:      1,
      ActiveCharacter:  chars[0],
      Done:             []Character{},
      Pending:          p,
   }
}
