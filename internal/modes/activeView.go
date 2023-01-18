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

   // "waelder/internal/language"
   ds "waelder/internal/datastructures"
   "waelder/internal/views"
   "waelder/internal/config"
)


// View method
func ActiveView(
   d ds.Data,
   windowHeight int,
   windowWidth int,
) []string{
   var res []string = make([]string, windowHeight)
   var runningIndex int = len(res) - 1

   // Generate String representing the current round
   s :=  formatBlock(
      d.CombatLog.Current.Done,
      d.CombatLog.Current.ActiveCharacter,
      d.CombatLog.Current.Pending,
      windowWidth,
      d.CombatLog.Current.RoundNumber,
      true,
   ) 

   // Uncomment for bar under layout
   // res[runningIndex] = strings.Repeat("─", windowWidth)
   // runningIndex--

   for i := len(s) - 1; i >= 0; i-- {
      res[runningIndex] = s[i]
      runningIndex--
   }

   for i:= len(d.CombatLog.PreviousRounds) - 1; i >= 0 && runningIndex >= 0; i-- {
      r := d.CombatLog.PreviousRounds[i]

      b := formatBlockOld(r.TurnSequence, windowWidth, r.RoundNumber, false)
      for j := len(b) - 1; j >= 0 && runningIndex >= 0; j-- {
         res[runningIndex] = b[j]
         runningIndex--

      }
   }

   return res
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
   }, separator)
}

func formatBlock(
   done           []ds.Character,
   active         ds.Character,
   pending        []ds.Character,
   windowWidth    int,
   roundNumber    int,
   isCurrentRound bool,
) []string {
   var s []string = make([]string, len(done) + len(pending) + 1 + 1)
   runningIndex := 1
 
   overlineLabel := fmt.Sprintf(" Round %d ", roundNumber)

   s[0]= "─┤" + config.StyleDarkRedBg.Render(overlineLabel) + "├"
   s[0] += strings.Repeat("─", windowWidth - (len(overlineLabel) + 3)) 

   for i := 0; i < len(done); i++ {
      s[runningIndex] = drawLine(done[i], false)
      runningIndex += 1
   }

   s[runningIndex] = drawLine(active, true);
   runningIndex += 1

   for i := 0; i < len(pending); i++ {
      s[runningIndex] = drawLine(pending[i], false)
      runningIndex += 1
   }

   return s
}
func formatBlockOld(
   done           []ds.Character,
   windowWidth    int,
   roundNumber    int,
   isCurrentRound bool,
) []string  {
   var s []string = make([]string, len(done) + 1)
   // Iterate over our choices
   
   overlineLabel := fmt.Sprintf("-┤ Round %d ├", roundNumber)
   s[0] = overlineLabel + strings.Repeat("─", windowWidth - len(overlineLabel) + 4)

   for i := 0; i < len(done); i++ {
      s[i+1] = drawLine(done[i], false);
   }
   return s
}

func ActiveUpdate(data *ds.Data, msg tea.KeyMsg) {
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
