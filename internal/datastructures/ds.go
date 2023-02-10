// ======================================================================
// Author: meisto
// Creation Date: Sun 15 Jan 2023 11:45:05 PM CET
// Description: -
// ======================================================================
package datastructures

import (
   "sort"
   "fmt"
   "log"
)

// Root data struct
type Data struct {
	Players  []string
	Allies   []string
	Enemies  []string
	Neutrals []string

	CombatLog   CombatLog
   characters  map[string]Character
   Graph       Graph
}

func GetData() Data { 
   return Data {
      Players: []string{},
      Allies:  []string{},
      Enemies: []string{},
      Neutrals:[]string{},

      CombatLog: CombatLog{
         PreviousRounds:   []Round{},
         Current: Round{
            RoundNumber:   -1,
            Done:             []string{},
            ActiveCharacter:  "",
            Pending:          []string{},
            Actions:          []Action{},
         },
      },
      characters: make(map[string]Character),
      Graph:      GetGraph(),
   } 
}
func (d *Data) AddPlayer(c Character) {
   if c.Affiliation != Player { 
      log.Fatal(c.Name, " is a ", c.Affiliation, " not a Player.")
   }
   d.Players = append(d.Players, c.Name)
   d.characters[c.Name] = c
}
func (d *Data) AddAlly(c Character) {
   if c.Affiliation != Ally { 
      log.Fatal(c.Name, " is a ", c.Affiliation, " not an Ally.")
   }
   d.Allies = append(d.Allies, c.Name)
   d.characters[c.Name] = c
}
func (d *Data) AddEnemy(c Character) {
   if c.Affiliation != Enemy { 
      log.Fatal(c.Name, " is a ", c.Affiliation, " not a Enemy.")
   }
   d.Enemies = append(d.Enemies, c.Name)
   d.characters[c.Name] = c
}
func (d *Data) AddNeutral(c Character) {
   if c.Affiliation != Neutral { 
      log.Fatal(c.Name, " is a ", c.Affiliation, " not a Neutral.")
   }
   d.Neutrals = append(d.Neutrals, c.Name)
   d.characters[c.Name] = c
}

func (data *Data) apply(a Action) {
   data.CombatLog.Current.Actions = 
      append(data.CombatLog.Current.Actions, a)

   for _, i := range(a.GetTargets()) {
      c := data.characters[i]

      data.characters[i] = a.Apply(c)
   }
}

func (data Data) GetCharacter(name string) Character {
   return data.characters[name]
}
func (data *Data) PrepareNextRound() {

   if !data.CombatLog.Current.IsDone() { return }

   rn := data.CombatLog.Current.RoundNumber

   // Only add current round to log if it was actually initialized
   if rn >= 0 {
      data.CombatLog.PreviousRounds = 
         append(data.CombatLog.PreviousRounds, data.CombatLog.Current)
   }

   pending := data.Players
   pending = append(pending, data.Allies...)
   pending = append(pending, data.Enemies...)
   pending = append(pending, data.Neutrals...)


	// Sort by initiative
	sort.Slice(pending, func(i, j int) bool {
		return data.characters[pending[i]].Stats.Initiative > 
         data.characters[pending[i]].Stats.Initiative
	})


   // Log
   log.Print(fmt.Sprintf("Prepared round number %d", rn + 1))


   // Jump to first character that is not already dead
   var i = 0
   for ; i < len(pending); i ++ {
      if !data.GetCharacter(pending[i]).IsDead() {break}
   }


   // Create new current round
	data.CombatLog.Current = Round{
		RoundNumber:      rn + 1,
      Done:             pending[:i],
		ActiveCharacter:  pending[i],
		Pending:          pending[i+1:],
      Actions:          []Action{},
   }
}

/*
*

	Progress the round.

	Returns true if a state transition was done, false if round is done

*
*/
func (data *Data) Step(ac Action) {
   data.apply(ac)

   data.StepWoAction()
}

func (data *Data) StepWoAction() {
   round := &data.CombatLog.Current

   
   // Give the actice previously active character its reaction back
   ch := data.GetCharacter(round.ActiveCharacter)
   ch.Stats.HasReaction = true

	if round.IsDone() {
      // Round is done, prepare next round and return
		data.PrepareNextRound()
      return
	}


	round.Done = append(round.Done, round.ActiveCharacter)

	// Sort remaining characters by initiative 
	sort.Slice(round.Pending, func(i, j int) bool {
		return data.characters[round.Pending[i]].Stats.Initiative > 
         data.characters[round.Pending[i]].Stats.Initiative
	})

	round.ActiveCharacter = round.Pending[0]
	round.Pending = round.Pending[1:]

   // Recursive call if next character is already dead
   if data.GetCharacter(data.CombatLog.Current.ActiveCharacter).IsDead() {
      data.StepWoAction()
   }
}

func (data Data) IsCombatOver() bool {
   { // Check if all enemies are dead
      isAlive := false
      for i := 0; i < len(data.Enemies); i ++ {
         ch := data.GetCharacter(data.Enemies[i])
         if !ch.IsDead() {
            isAlive = true
            break
         }
      }
      if !isAlive { return true }  
   }

   return false
}
