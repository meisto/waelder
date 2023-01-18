// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 21 Nov 2022 05:08:56 PM CET
// Description: -
// ======================================================================
package modes

import(
   "fmt"
   "strings"

   tea "github.com/charmbracelet/bubbletea"

   // "dntui/internal/language"
   ds "dntui/internal/datastructures"
   "dntui/internal/views"
   "dntui/internal/config"
)


// View method
func activeView(
   d ds.Data,
   windowHeight int,
   windowWidth int,
) string {
   // Generate String representing the current round
   s :=  formatBlock(
      d.CombatLog.Current.Done,
      d.CombatLog.Current.ActiveCharacter,
      d.CombatLog.Current.Pending,
      windowWidth,
      d.CombatLog.Current.RoundNumber,
      true,
   ) 

   // Print trace of previousrounds
   linecount := strings.Count(s, "\n") + 1 // + 1 for bottom line
   for i:= len(d.CombatLog.PreviousRounds) - 1; i >= 0 && linecount <= windowHeight; i-- {
      r := d.CombatLog.PreviousRounds[i]

      b := formatBlockOld(r.TurnSequence, windowWidth, r.RoundNumber, false)
      linecount += strings.Count(b, "\n")

      s = b + s
   }
   s += strings.Repeat("─", windowWidth) + "\n"





   return s
}


func drawLine(
   char ds.Character,
   isActive bool,
) string {
   initiative := fmt.Sprintf("👞%2d", char.Stats.Initiative)

   charNameWidth := 20
   charName := char.Name
   if len(charName) > charNameWidth {
      charName = charName[:charNameWidth-3] + "..."
   }  else {
      charName = fmt.Sprintf(" %17s ",charName) // not selected
   }


   // Generate 
   hpPercentage := float64(char.Stats.Hp) / float64(char.Stats.Max_hp)
   
   isDead := hpPercentage == 0.0

   f := func(s string) string {return views.FormatString(s, isActive, isDead)}

   health := views.FormatHealthString(10, hpPercentage, isActive)
   healthNumeral := f(fmt.Sprintf("%03d/%03d", char.Stats.Hp, char.Stats.Max_hp))

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
   done           []ds.Character,
   active         ds.Character,
   pending        []ds.Character,
   windowWidth    int,
   roundNumber    int,
   isCurrentRound bool,
) string {
   var s string
   // Iterate over our choices
   
   prefix := "─┤" + config.StyleDarkRedBg.Render(fmt.Sprintf(" Round %d ", roundNumber)) + "├"

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
func formatBlockOld(
   done           []ds.Character,
   windowWidth    int,
   roundNumber    int,
   isCurrentRound bool,
) string  {
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

func activeUpdate(data *ds.Data, msg tea.KeyMsg) {
   switch msg.String(){
      case "n":

         if data.CombatLog.Current.IsDone() {
            // Trigger start of new Round
            oldRound, newRound := data.CombatLog.Current.GetNextRound()

            data.CombatLog.PreviousRounds   = append(data.CombatLog.PreviousRounds, oldRound)
            data.CombatLog.Current          = newRound
         } else {
            data.CombatLog.Current.Step()
         }
   }
}
