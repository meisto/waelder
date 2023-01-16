// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Sun 15 Jan 2023 11:45:05 PM CET
// Description: -
// ======================================================================

package model

import (
   "sort"

   c "dntui/internal/model/charactermodel"
)

type Log struct {
   initialized    bool
   PreviousRounds []Round
   Current        CurrentRound
}

type Round struct {
   RoundNumber    int
   TurnSequence   []c.Character
}

type CurrentRound struct {
   RoundNumber       int
   ActiveCharacter   c.Character
   Done              []c.Character
   Pending           []c.Character
}

func (round CurrentRound) IsDone() bool { return len(round.Pending) == 0 }

/**
   Progress the round. 

   Returns true if a state transition was done, false if round is done
**/
func (round CurrentRound) Step() bool{

   // Break
   if round.IsDone() {return false}
   
   round.Done = append(round.Done, round.ActiveCharacter)

   // Sort by initiative
   comp := func(i,j int) bool {return round.Pending[i].Initiative > round.Pending[j].Initiative}
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
      TurnSequence:  round.Done,
   }

   p := round.Pending

   // Sort by initiative
   sort.Slice(p, func(i,j int) bool {
      return round.Pending[i].Initiative > round.Pending[j].Initiative
   })

   b := CurrentRound {
      RoundNumber:      round.RoundNumber + 1,
      ActiveCharacter:  p[0],
      Done:             []c.Character{},
      Pending:          p[1:],
   }

   return a, b
}

func CreateRound(chars []c.Character) CurrentRound {
   // Sort by initiative
   sort.Slice(chars, func(i,j int) bool {
      return chars[i].Initiative > chars[j].Initiative
   })
   

   return CurrentRound {
      RoundNumber:      1,
      ActiveCharacter:  chars[0],
      Done:             []c.Character{},
      Pending:          chars[1:],
   }
}
