// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Wed 25 Jan 2023 05:40:23 PM CET
// Description: -
// ======================================================================
package root

import (
   "fmt"
   "strconv"

   "github.com/muesli/termenv"

	ds "waelder/internal/datastructures"
	"waelder/internal/datastructures/graph"
	"waelder/internal/layouts"
   "waelder/internal/modes"
   "waelder/internal/renderer"
)

func addActionPopupSequence(
   output   *termenv.Output,
   data     *ds.Data,
   tg       *graph.Graph,
   n0       graph.Node,
   nLast    graph.Node,
   x        int,
   y        int,
   layout   layouts.Layout,

) {
   // Get Intermediate nodes
   n1 := graph.GetNode()

   c1Header := "Whom should %s attack?"
   c1Footer := "FOOTER"


   var c1 modes.Choice
   { // Generate first popup menu
      a := make([]string, len(data.Players))
      b := make([]string, len(data.Players))
      for i := 0; i < len(data.Players); i++ {
         b[i] = " x "
         a[i] = data.Players[i]
      }
      c1 = modes.GetChoice(
         0,
         fmt.Sprintf(c1Header, data.CombatLog.Current.ActiveCharacter),
         c1Footer,
         b,
         a,
      )
   }
 
   w := 50
   h := len(data.Players) + 4


   redrawChoicePopup := func() {
      layouts.PopUp(output, c1.ToRenderField(x,y,w) , x, y, w, h)
   }


   // Add Graph traversal edges
	tg.AddEdge(
		graph.GetEdge(
			n0,
			"p",
			n1,
         func() {
            c1.Header = fmt.Sprintf(c1Header, data.CombatLog.Current.ActiveCharacter)
            redrawChoicePopup()
         },
		),
	)

   // First popup
   f := func(s string, g func()) {
      tg.AddEdge(graph.GetEdge(n1, s, n1, g))
   }

   f("n", func() {c1.Step(); redrawChoicePopup() })
   f("0", func() {c1.GoTo(9); redrawChoicePopup() })
   f("1", func() {c1.GoTo(0); redrawChoicePopup() })
   f("2", func() {c1.GoTo(1); redrawChoicePopup() })
   f("3", func() {c1.GoTo(2); redrawChoicePopup() })
   f("4", func() {c1.GoTo(3); redrawChoicePopup() })
   f("5", func() {c1.GoTo(4); redrawChoicePopup() })
   f("6", func() {c1.GoTo(5); redrawChoicePopup() })
   f("7", func() {c1.GoTo(6); redrawChoicePopup() })
   f("8", func() {c1.GoTo(7); redrawChoicePopup() })
   f("9", func() {c1.GoTo(8); redrawChoicePopup() })




   tg.AddEdge(
      graph.GetEdge(
         n1,
         "<ENTER>",
         nLast,
         func() {
            print(c1.GetSelection())
            content := renderer.GenerateField(
               []renderer.RenderLine{
                  renderer.GenerateLine(w - 2, []renderer.Renderable{renderer.GenerateNoRenderNode("123")}),
               },
            )

            rl := layouts.ReadLinePopUp(output, content, x, y, w, h,)
            n, err := strconv.Atoi(rl)
            for err != nil || n < 0{
               content = renderer.GenerateField(
                  []renderer.RenderLine{
                     renderer.GenerateLine(w - 2, []renderer.Renderable{renderer.GenerateNoRenderNode("123ERROR")}),
                  },
               )
               rl := layouts.ReadLinePopUp(output, content, x, y, w, h,)
               n, err = strconv.Atoi(rl)

               

            }
            
            at := ds.Attack{
               Turn: data.CombatLog.Current.RoundNumber,
               Source: data.CombatLog.Current.ActiveCharacter,
               Targets: []string{c1.GetSelection()},
               HasHit: true,
               Damage: n,
               Range: ds.Meele,
            }
            data.AddAction(at)


            print(n)

            layout.Reset(output, *data)
            
         },
      ),
   )
}
