// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 21 Nov 2022 05:08:56 PM CET
// Description: -
// ======================================================================
package modes

import (
	"fmt"
	"strings"
   "log"

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
	x int,
	y int,
	windowHeight int,
	windowWidth int,
) {
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


	l := windowHeight
	a.RenderBlock(output, x, y, l, true, 100000)
	l -= a.Length()



	i := roundNumber - 1
	for ; i >= 0 && l >= 0; i-- {

      // TODO Update chars
      b := previousBlocks[i]
		b.RenderBlock(output, x, y, l, true, 10000)
		l -= b.GetHeight()
	}

   

   if d.CombatLog.Current.IsDone() && lastBlockSave < roundNumber {
      lastBlockSave = roundNumber
      log.Print("isdone")
      previousBlocks = append(
         previousBlocks,
         formatBlock(chars, -1, windowWidth, roundNumber, false),
      )

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
	} else {
		charName = fmt.Sprintf(" %17s ", charName) // not selected
	}


   hp := char.Stats.Hp
   if hp < 0 { hp = 0 }

	// Generate
	hpPercentage := float64(hp) / float64(char.Stats.Max_hp)
   if char.Stats.Max_hp == 0 { hpPercentage = 0.0}
   if hpPercentage < 0 { hpPercentage = 0.0 }

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

	health := FormatHealthString(10, hpPercentage, isActive, heartIcon)
	healthNumeral := f(fmt.Sprintf("%03d/%03d", hp, char.Stats.Max_hp))

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
      b = renderer.GenerateNode(fmt.Sprintf(" Round %d ", roundNumber), "darkRedBg")
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
