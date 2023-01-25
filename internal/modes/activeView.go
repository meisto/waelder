// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 21 Nov 2022 05:08:56 PM CET
// Description: -
// ======================================================================
package modes

import(
   "fmt"
   "strings"

   "github.com/muesli/termenv"

   // "waelder/internal/language"
   ds "waelder/internal/datastructures"
   "waelder/internal/renderer"
)


// View method
func ActiveView(
   output *termenv.Output,
   d ds.Data,
   x int,
   y int,
   windowHeight int,
   windowWidth int,
) {

   // Generate String representing the current round
   a :=  formatBlock(
      d.CombatLog.Current.Done,
      d.CombatLog.Current.ActiveCharacter,
      d.CombatLog.Current.Pending,
      windowWidth,
      d.CombatLog.Current.RoundNumber,
      true,
   ) 

   l := windowHeight
   a.RenderBlock(output, x, y, l, true, true)

   l -= a.Length()

   i := len(d.CombatLog.PreviousRounds) - 1
   for ; i >= 0 && l >= 0; i-- {
      r := d.CombatLog.PreviousRounds[i]

      b := formatBlockOld(r.TurnSequence, windowWidth, r.RoundNumber, false)
      b.RenderBlock(output, x, y, l, true, true)
      l -= b.GetHeight()
   }
}


func drawLine(
   char ds.Character,
   isActive bool,
   width int,
) renderer.RenderLine {
   // Sonderzeichen: 👞
   // "♡♥❤ ",
   // heartIcon := "♥" // "♡"
   
   // Hardcoded settings
   sprintIcon := "👞"
   charNameWidth := 20
   // heartIcon := "X"
   heartIcon := "♥" // "♡"

   initiative := fmt.Sprintf("%s %2d", sprintIcon, char.Stats.Initiative)

   charName := char.Name
   if len(charName) > charNameWidth {
      charName = charName[:charNameWidth-3] + "..."
   }  else {
      charName = fmt.Sprintf(" %17s ",charName) // not selected
   }

   // Generate 
   hpPercentage := float64(char.Stats.Hp) / float64(char.Stats.Max_hp)
   isDead := hpPercentage == 0.0

   f := func(s string) renderer.RenderNode {

      var style string
      if isActive {
         style = "selected"
      } else if isDead {
         style = "dead"
      } else {
         style = "default"
      }
      return renderer.GenerateNode(s, style)
   }

   health := views.FormatHealthString(10, hpPercentage, isActive, heartIcon)
   healthNumeral := f(fmt.Sprintf("%03d/%03d", char.Stats.Hp, char.Stats.Max_hp))

   // Generate separator between row entries
   separator := f(" ")

   // Assebmle string
   a := renderer.GenerateLine(
      width, 
      []renderer.Renderable{
         f(initiative),
         separator,
         f(charName),
         separator,
         health,
         separator,
         healthNumeral,
         separator,
         f("RK 1011"),
      },
   )

   // Pad/cut length, 
   /*
   if len(a) > width { a = a[:width - 3] + "..." }
   */


   return a
}

func formatBlock(
   done           []ds.Character,
   active         ds.Character,
   pending        []ds.Character,
   windowWidth    int,
   roundNumber    int,
   isCurrentRound bool,
) renderer.RenderField {
   var s []renderer.RenderLine =
      make([]renderer.RenderLine, len(done) + 1 + len(pending) + 1)
   runningIndex := 1
 

   a := renderer.GenerateNoRenderNode("──┤")
   b := renderer.GenerateNode(fmt.Sprintf(" Round %d ", roundNumber), "darkRedBg")
   c := renderer.GenerateNoRenderNode("├─")
   d := windowWidth - a.Length() - b.Length() - c.Length()

   s[0] = renderer.GenerateLine(
      windowWidth,
      []renderer.Renderable {
         a,
         b,
         c,
         renderer.GenerateNoRenderNode(strings.Repeat("─", d)),
      },
   )

   // ─┤├

   for i := 0; i < len(done); i++ {
      s[runningIndex] = drawLine(done[i], false, windowWidth)
      runningIndex += 1
   }

   s[runningIndex] = drawLine(active, true, windowWidth);
   runningIndex += 1

   for i := 0; i < len(pending); i++ {
      s[runningIndex] = drawLine(pending[i], false, windowWidth)
      runningIndex += 1
   }

   return renderer.GenerateField(s)
}

func formatBlockOld(
   done           []ds.Character,
   windowWidth    int,
   roundNumber    int,
   isCurrentRound bool,
) renderer.RenderField  {
   var s []renderer.RenderLine =
      make([]renderer.RenderLine, len(done) + 1)
   runningIndex := 1
 
   a := renderer.GenerateNoRenderNode("──┤")
   b := renderer.GenerateNoRenderNode(fmt.Sprintf(" Round %d ", roundNumber))
   c := renderer.GenerateNoRenderNode("├─")
   d := windowWidth - a.Length() - b.Length() - c.Length()

   s[0] = renderer.GenerateLine(
      windowWidth,
      []renderer.Renderable {
         a,
         b,
         c,
         renderer.GenerateNoRenderNode(strings.Repeat("─", d)),
      },
   )

   for i := 0; i < len(done); i++ {
      s[runningIndex] = drawLine(done[i], false, windowWidth)
      runningIndex += 1
   }

   return renderer.GenerateField(s)
}

func ActiveUpdate(data *ds.Data, msg string) {
   switch msg {
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
