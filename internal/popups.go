// ======================================================================
// Author: meisto
// Creation Date: Wed 25 Jan 2023 05:40:23 PM CET
// Description: -
// ======================================================================
package root

import (
   "fmt"
   "strconv"

   "github.com/muesli/termenv"

	ds "waelder/internal/datastructures"
	"waelder/internal/layouts"
   "waelder/internal/modes"
   "waelder/internal/renderer"
   "waelder/internal/asciiart"
)

func addActionPopupSequence(
   output   *termenv.Output,
   data     *ds.Data,
   n0       ds.Node,
   nLast    ds.Node,
   x        int,
   y        int,
   width    int,
   layout   layouts.Layout,
) {
   // Get Intermediate nodes
   n1 := ds.GetNode()
   n2 := ds.GetNode()
   n3 := ds.GetNode()

   tg := &(data.Graph)

   c1Header := " Whom should %s attack?"


   // Choice module to pick target
   var c1 modes.Choice
   labels  := append(data.Enemies, data.Neutrals...)
   labels   = append(labels, data.Allies...)
   labels   = append(labels, data.Players...)
   { // Generate first popup menu
      l := len(data.Players) + len(data.Allies) + len(data.Enemies) +
         len(data.Neutrals)
      buttons := []string{
         "1", "2", "3", "4", "5", "6","7","8","9","0","q","w","e","r","t","z","u","i","l","y"}
      for i := 0; i < len(buttons); i++ {buttons[i] = fmt.Sprintf(" %s ", buttons[i])}
      styledLabels := make([]renderer.Renderable, len(labels))

      f := func(x renderer.Renderable) renderer.Renderable {
         return renderer.GenerateLine(x.Length() + 1,[]renderer.Renderable{renderer.GenerateNoRenderNode(" "), x})
      }

      i := 0
      for j := 0; j < len(data.Enemies); j++ {
         ch := data.GetCharacter(data.Enemies[j])
         styledLabels[i] = f(renderer.GenerateNode(ch.Name, ch.GetStyle()))
         i += 1
      }
      for j := 0; j < len(data.Neutrals); j++ {
         ch := data.GetCharacter(data.Neutrals[j])
         styledLabels[i] = f(renderer.GenerateNode(ch.Name, ch.GetStyle()))
         i += 1
      }
      for j := 0; j < len(data.Allies); j ++ {
         ch := data.GetCharacter(data.Allies[j])
         styledLabels[i] = f(renderer.GenerateNode(ch.Name, ch.GetStyle()))
         i += 1
      }
      for j := 0; j < len(data.Players); j++ {
         ch := data.GetCharacter(data.Players[j])
         styledLabels[i] = f(renderer.GenerateNode(ch.Name, ch.GetStyle()))
         i += 1
      }

      c1 = modes.GetChoice(
         fmt.Sprintf(c1Header, data.CombatLog.Current.ActiveCharacter),
         buttons[:l],
         styledLabels,
      )
   }


   // Choice module to pick target
   var c2 modes.Choice
   {
      b := []string{" 1 "," 2 "," 3 "," 4 "}
      s1 := func( s string) renderer.Renderable{ return renderer.GenerateNode(s, "italic")}
      s2 := func( s string) renderer.Renderable{ return renderer.GenerateNode(s, "bold")}
      labels := []renderer.Renderable{
         renderer.GenerateLine(width - 5, []renderer.Renderable{s1(" Meele    "), s2(asciiart.OneLineSword)}),
         renderer.GenerateLine(width - 5, []renderer.Renderable{s1(" Ranged   "), s2(asciiart.OneLineArrow)}),
         renderer.GenerateLine(width - 5, []renderer.Renderable{s1(" Magical  "), s2(asciiart.OneLineFire)}),
         renderer.GenerateLine(width - 5, []renderer.Renderable{s1(" Healing  "), s2(asciiart.OneLineFire)}),
      }

      c2 = modes.GetChoice("Which mode of attack?", b, labels)
   }

   var c3 modes.Choice
   {
      b := []string{" Yes ", " No "}
      labels := []renderer.Renderable {
         renderer.GenerateNoRenderNode(""),
         renderer.GenerateNoRenderNode(""),
      }

      c3 = modes.GetChoice("Did they hit?", b, labels)
   }
 
   redrawChoicePopup := func(c modes.Choice) {
      layouts.PopUp(output, c.ToRenderField(x,y,width) , x, y, *data)
   }


   // Add Graph traversal edges
	tg.AddEdge(ds.GetEdge( n0, "A", n1, func() {
         c1.Header = fmt.Sprintf(c1Header, data.CombatLog.Current.ActiveCharacter)
         redrawChoicePopup(c1)
      },
      "Attack",
   ))


   // First popup control
   {
      f := func(s string, g func(), desc string) { 
         tg.AddEdge(ds.GetEdge(n1, s, n1, g, desc)) 
      }

      f("j", func() {c1.Forward();  redrawChoicePopup(c1)}, "Next")
      f("k", func() {c1.Backward();  redrawChoicePopup(c1)}, "Previous")
      f("0", func() {c1.GoTo(9); redrawChoicePopup(c1)}, "Select 0")
      f("1", func() {c1.GoTo(0); redrawChoicePopup(c1)}, "Select 1")
      f("2", func() {c1.GoTo(1); redrawChoicePopup(c1)}, "Select 2")
      f("3", func() {c1.GoTo(2); redrawChoicePopup(c1)}, "Select 3")
      f("4", func() {c1.GoTo(3); redrawChoicePopup(c1)}, "Select 4")
      f("5", func() {c1.GoTo(4); redrawChoicePopup(c1)}, "Select 5")
      f("6", func() {c1.GoTo(5); redrawChoicePopup(c1)}, "Select 6")
      f("7", func() {c1.GoTo(6); redrawChoicePopup(c1)}, "Select 7")
      f("8", func() {c1.GoTo(7); redrawChoicePopup(c1)}, "Select 8")
      f("9", func() {c1.GoTo(8); redrawChoicePopup(c1)}, "Select 9")

      tg.AddEdge(ds.GetEdge(n1, "<ENTER>", n2, func(){
         layout.Reset(output, *data)
         redrawChoicePopup(c2)
      }, "Confirm"))
   }


   // Second popup control
   {
      f := func(s string, g func(), desc string) { 
         tg.AddEdge(ds.GetEdge(n2, s, n2, g, desc)) 
      }
      f("j", func() {c2.Forward(); redrawChoicePopup(c2)}, "Next")
      f("k", func() {c2.Backward(); redrawChoicePopup(c2)}, "Previous")
      f("1", func() {c2.GoTo(0); redrawChoicePopup(c2)}, "Meele")
      f("2", func() {c2.GoTo(1); redrawChoicePopup(c2)}, "Ranged")
      f("3", func() {c2.GoTo(2); redrawChoicePopup(c2)}, "Magic")
      f("4", func() {c2.GoTo(3); redrawChoicePopup(c2)}, "Healing")

      tg.AddEdge(ds.GetEdge(n2, "<ENTER>", n3, func(){
         layout.Reset(output, *data)
         redrawChoicePopup(c3)
      }, "tmp"))
   }


   // Third popup control
   {
      f := func(s string, g func(), desc string) { 
         tg.AddEdge(ds.GetEdge(n3, s, n3, g, desc)) 
      }

      f("j", func() {c3.Forward(); redrawChoicePopup(c3)}, "Next")
      f("k", func() {c3.Backward(); redrawChoicePopup(c3)}, "Previous")
      f("1", func() {c3.GoTo(0); redrawChoicePopup(c3)}, "Yes")
      f("2", func() {c3.GoTo(1); redrawChoicePopup(c3)}, "No")

   }


   tg.AddEdge(
      ds.GetEdge(
         n3,
         "<ENTER>",
         nLast,
         func() {
            layout.Reset(output, *data)


            var n int = 0

            // Prompt for damage if not hit
            if c3.GetIndex() == 0 {
               prompt := renderer.GenerateNoRenderNode("How much damage did they deal?")   
               content := renderer.GenerateField(
                  []renderer.RenderLine{
                     renderer.GenerateLineFromOne(width - 2, prompt),
                  },
               )
               rl := layouts.ReadLinePopUp(output, content, x, y, *data)

               // Try to cast string to int
               n2, err := strconv.Atoi(rl)
               n = n2
               for err != nil || n < 0 || n > 1000 {
                  prompt = renderer.GenerateNoRenderNode("Illegal input, please try again. ")

                  content = renderer.GenerateField(
                     []renderer.RenderLine{
                        renderer.GenerateLineFromOne(
                           width - 2,
                           prompt,
                        ),
                     },
                  )
                  rl := layouts.ReadLinePopUp(output, content, x, y, *data)
                  n, err = strconv.Atoi(rl)
               }
            } 
            
            at := ds.Attack{
               Round: data.CombatLog.Current.RoundNumber,
               Source: data.CombatLog.Current.ActiveCharacter,
               Targets: []string{labels[c1.GetIndex()]},
               HasHit: c3.GetIndex() == 0,
               Damage: n,
               Range: ds.ToAttackType(c3.GetIndex()),
            }

            // AAAAAAAA
            data.Step(at)



            layout.Reset(output, *data)
            
         },
      "tmp"),
   )
}
