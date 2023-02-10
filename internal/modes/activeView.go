// ======================================================================
// Author: meisto
// Creation Date: Mon 21 Nov 2022 05:08:56 PM CET
// Description: -
// ======================================================================
package modes

import (
	"fmt"
	"strings"

	"github.com/muesli/termenv"

	// "waelder/internal/language"
	ds "waelder/internal/datastructures"
	"waelder/internal/renderer"
)

// Track previous round as to not need to recalculate
var previousBlocks []renderer.RenderField = []renderer.RenderField{}
var lastBlockSave int = -1

// View method
func activeView(
	output *termenv.Output,
	d ds.Data,
	windowHeight int,
	windowWidth int,
) renderer.RenderField {
   chars := make([]ds.Character, d.CombatLog.Current.NumberCombatants())
   activeIndex := len(d.CombatLog.Current.Done)
   roundNumber := d.CombatLog.Current.RoundNumber

   for i := 0; i < len(chars); i++ {
      if i < activeIndex {
         chars[i] = d.GetCharacter(d.CombatLog.Current.Done[i])
      } else if i == activeIndex {
         chars[i] = d.GetCharacter(d.CombatLog.Current.ActiveCharacter)
      } else {
         chars[i] = d.GetCharacter(d.CombatLog.Current.Pending[i - activeIndex - 1])
      }
   }
   
	// Generate String representing the current round
	a := formatBlock(
      chars,
      activeIndex,
		windowWidth,
		roundNumber,
		true,
	)

   /*
      l := windowHeight
      l -= a.Length()

      i := roundNumber - 1
      for ; i >= 0 && l >= 0 && i < len(previousBlocks); i-- {

         // TODO Update chars
         b := previousBlocks[i]
         l -= b.GetHeight()
         a = b.Join(a)
      }

      

      if d.CombatLog.Current.IsDone() && lastBlockSave < roundNumber {
         lastBlockSave = roundNumber
         previousBlocks = append(
            previousBlocks,
            formatBlock(chars, -1, windowWidth, roundNumber, false),
         )
      
      }
   */
   return a

}

func drawLine(
	char ds.Character,
	isActive bool,
	width int,
) renderer.RenderLine {

	// Hardcoded settings
	sprintIcon     := "👞"
	heartIcon      := "♥" // "♡♥❤ ",
   armourIcon     := ""
   reactionIcon   := ""
	charNameWidth := 20


   // Format character name
	var charName renderer.Renderable
   {
      name := char.Name

      if len(name) > charNameWidth {
         name = name[:charNameWidth-3] + "..."
      } else {
         name = fmt.Sprintf(" %17s ", name) // not selected
      }

      var style string
      if !isActive {
         switch char.Affiliation {
            case ds.Ally: style     = "allyUnit"
            case ds.Player: style   = "playerUnit"
            case ds.Neutral: style  = "neutralUnit"
            case ds.Enemy: style    = "enemyUnit"
         }
      } else {
         switch char.Affiliation {
            case ds.Ally: style     = "allyUnitActive"
            case ds.Player: style   = "playerUnitActive"
            case ds.Neutral: style  = "neutralUnitActive"
            case ds.Enemy: style    = "enemyUnitActive"
         }

      }

      charName = renderer.GenerateNode(name, style)
   }


	// Format HP info
   hp := char.Stats.Hp
   if hp < 0 { hp = 0 }
	hpPercentage := float64(hp) / float64(char.Stats.Max_hp)

   // Capture special cases for mathematical stability
   if char.Stats.Max_hp == 0 { hpPercentage = 0.0}
   if hpPercentage < 0 { hpPercentage = 0.0 }

   // Reaction
   var reaction string
   if char.Stats.HasReaction {
      reaction = "  "
   } else {
      reaction = fmt.Sprintf("%s ", reactionIcon)
   }


   // Format other stuff
	initiative := fmt.Sprintf("%s %2d", sprintIcon, char.Stats.Initiative)
	health := FormatHealthString(10, hpPercentage, isActive, heartIcon)
	healthNumeral := fmt.Sprintf("%03d/%03d", hp, char.Stats.Max_hp)
   rk := fmt.Sprintf("%s %03d", armourIcon, char.Stats.Armour)


   // Function to style elements
	f := func(s string) renderer.RenderNode {
	   isDead := hpPercentage == 0.0
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

	// Generate separator between row entries
	separator := f("  ")

	// Assebmle string
	a := renderer.GenerateLine(
		width,
		[]renderer.Renderable{
			f(initiative),
			separator,
			charName,
			separator,
			health,
			separator,
			f(healthNumeral),
			separator,
			f(rk),
			separator,
         f(reaction),
		},
	)

	// Pad/cut length,
	/*
	   if len(a) > width { a = a[:width - 3] + "..." }
	*/

	return a
}

func formatBlock(
	characters  []ds.Character,
   activeIndex int,
	windowWidth int,
	roundNumber int,
	isCurrentRound bool,
) renderer.RenderField {
	var s []renderer.RenderLine = make([]renderer.RenderLine, len(characters) + 1)
	runningIndex := 1

	a := renderer.GenerateNoRenderNode("──┤")
   var b renderer.Renderable
   if isCurrentRound {
      b = renderer.GenerateNode(fmt.Sprintf(" Round %d ", roundNumber + 1), "darkRedBg")
   } else {
      b = renderer.GenerateNoRenderNode(fmt.Sprintf(" Round %d ", roundNumber))
   }
	c := renderer.GenerateNoRenderNode("├─")
	d := windowWidth - a.Length() - b.Length() - c.Length()

	s[0] = renderer.GenerateLine(
		windowWidth,
		[]renderer.Renderable{
			a,
			b,
			c,
			renderer.GenerateNoRenderNode(strings.Repeat("─", d)),
		},
	)

	// ─┤├

	for i := 0; i < len(characters); i++ {
		s[runningIndex] = drawLine(characters[i], i == activeIndex, windowWidth)
		runningIndex += 1
	}

	return renderer.GenerateField(s)
}
