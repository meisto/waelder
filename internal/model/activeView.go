// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 21 Nov 2022 05:08:56 PM CET
// Description: -
// ======================================================================
package model

import(
   "fmt"
   "strings"

   tea "github.com/charmbracelet/bubbletea"

   // "dntui/internal/language"
   cm "dntui/internal/model/charactermodel"
   "dntui/internal/views"
   "dntui/internal/data"
)

var log Log = Log {
   initialized:      false,
   PreviousRounds:   []Round{},
}

// View method
func activeView(
   d data.Data,
   windowHeight int,
   windowWidth int,
) string {

   // Get current turn order when not set
   if !log.initialized{
      characters := []cm.Character{}
      

      // characters := append(d.Npcs, d.Pcs...)
      // characters  = append(characters, d.Others...)
   
      log.Current = CreateRound(characters)
      log.initialized = true

   }

   // Generate String representing the current round
   s :=  formatBlock(
      log.Current.Done,
      log.Current.ActiveCharacter,
      log.Current.Pending,
      windowWidth,
      log.Current.RoundNumber,
      true,
   ) 

   // Print trace of previousrounds
   linecount := strings.Count(s, "\n") + 1 // + 1 for bottom line
   for i:= len(log.PreviousRounds) - 1; i >= 0 && linecount <= windowHeight; i-- {
      r := log.PreviousRounds[i]

      b := formatBlock2(r.TurnSequence, windowWidth, r.RoundNumber, false)
      linecount += strings.Count(s, "\n")

      s = b + s

   }
   s += strings.Repeat("─", windowWidth) + "\n"





   return s
}


func drawLine(
   char cm.Character,
   isActive bool,
) string {
   initiative := fmt.Sprintf("👞%2d", char.Initiative)
   charName := fmt.Sprintf(" %10s ",char.Name) // not selected

   // Generate 
   hpPercentage := float64(char.Hp) / float64(char.Max_hp)
   isDead := hpPercentage == 0.0

   f := func(s string) string {return views.FormatString(s, isActive, isDead)}

   health := views.FormatHealthString(10, hpPercentage, isActive)
   healthNumeral := f(fmt.Sprintf("%03d/%03d", char.Hp, char.Max_hp))

   // Generate separator between row entries
   separator := f(" ")

   // Assebmle string
   return strings.Join([]string{
      f(initiative),
      f(charName),
      health,
      healthNumeral,
      f("🛡 1011"),
   },
   separator) + "\n"
}

func formatBlock(
   done           []cm.Character,
   active         cm.Character,
   pending        []cm.Character,
   windowWidth    int,
   roundNumber    int,
   isCurrentRound bool,
) string {
   var s string
   // Iterate over our choices
   
   prefix := fmt.Sprintf("─┤ Round %d ├", roundNumber)

   s += prefix + strings.Repeat("─", windowWidth - len(prefix)) 
   s += "\n"

   for i := 0; i < len(done); i++ {
      s += drawLine(done[i], false);
   }
   s += drawLine(active, true);
   for i := 0; i < len(pending); i++ {
      s += drawLine(pending[i], false);
   }
   return s
}
func formatBlock2(
   done           []cm.Character,
   windowWidth    int,
   roundNumber    int,
   isCurrentRound bool,
) string {
   var s string
   // Iterate over our choices
   
   prefix := fmt.Sprintf("─┤ Round %d ├", roundNumber)

   s += prefix + strings.Repeat("─", windowWidth - len(prefix)) 
   s += "\n"

   for i := 0; i < len(done); i++ {
      s += drawLine(done[i], false);
   }

   return s

}

func activeUpdate(m *model, msg tea.KeyMsg) {
   switch msg.String(){
      case "n":

         if log.Current.IsDone() {
            // Trigger start of new Round
            oldRound, newRound := log.Current.GetNextRound()


            log.PreviousRounds   = append(log.PreviousRounds, oldRound)
            log.Current          = newRound

         }
   }
}
