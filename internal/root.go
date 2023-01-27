// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Fri 20 Jan 2023 02:00:59 AM CET
// Description: -
// ======================================================================
package root

import (
   "log"
	"database/sql"
	"path"

	"github.com/muesli/termenv"

	"golang.org/x/term"

	ds "waelder/internal/datastructures"
	"waelder/internal/datastructures/graph"
	"waelder/internal/wio"
	"waelder/internal/layouts"
   "waelder/internal/modes"
	"waelder/internal/triggers"
)

func NewRun(dbHandle *sql.DB, elchan chan string, output *termenv.Output) {
   log.Print("Start Run")

	// TODO: Replace the defaults with a loading procedure
	root := "/home/tobias/Documents/code/golang/src/waelder"
   character1 := wio.LoadCharacterFromFile(path.Join(root, "data/char1.json"))
	character2 := wio.LoadCharacterFromFile(path.Join(root, "data/char2.json"))
	character3 := wio.LoadCharacterFromFile(path.Join(root, "data/char3.json"))
	character4 := wio.LoadCharacterFromFile(path.Join(root, "data/char4.json"))

	f := func(c ds.Character) ds.Character {
		return wio.SyncCharacterWithDatabase(dbHandle, c)
	}

   ch1 := f(character1)
	ch2 := f(character2)
	ch3 := f(character3)
	ch4 := f(character4)

	data := ds.GetData()

   data.AddPlayer(ch1)
   data.AddPlayer(ch2)
   data.AddPlayer(ch3)
   data.AddPlayer(ch4)
   data.PrepareNextRound()

	width, height, _ := term.GetSize(0)
	height -= 1

   // Set layout
	layout := layouts.TwoOneHorizontalSplit(height, width)


	var InitialNode graph.Node = graph.GetNode()

	tg := graph.GetGraph(InitialNode)
	tg.AddEdge(
		graph.GetEdge(
			InitialNode,
			"n",
			InitialNode,
			func() {
            data.Step()

				// Redraw fields with this mode
				layout.UpdateMode(output, data, modes.ActiveMode)
			},
		),
	)

	tg.AddEdge(
		graph.GetEdge(
			InitialNode,
			"o",
			InitialNode,
			func() {
				layout.Reset(output, data)
			},
		),
	)

   addActionPopupSequence(
      output,
      &data,
      &tg,
      InitialNode,
      InitialNode,
      20,
      20,
      layout,
   )

   n2 := graph.GetNode()
   tg.AddEdge(
      graph.GetEdge(InitialNode, "i", n2, func() {
         var f layouts.FieldInterface = layout.Fields[1]

         x := f.(layouts.ScrollField)
         y := &x
         y.SetBorder(layouts.DefaultBorderStyle.Style("darkGreenFg"))
         y.DrawBorder(output)
         y.DrawContent(output, data)
      }),
   )

   tg.AddEdge(graph.GetEdge(n2, "+", n2, func() {
               layouts.ScrollUp()
               log.Print(layouts.GetIndex())
               layout.Fields[1].DrawBorder(output)
   }))
   tg.AddEdge(graph.GetEdge(n2, "-", n2, func() {
      switch a := layout.Fields[1].(type) {
         case layouts.ScrollField:
               layouts.ScrollDown()
               log.Print(layouts.GetIndex())
               a.DrawBorder(output)
               a.DrawContent(output, data)
      }
   }))

   tg.AddEdge(
      graph.GetEdge(n2, "i", InitialNode, func() {
         var f layouts.FieldInterface = layout.Fields[1]

         switch x := f.(type) {
            case layouts.ScrollField:
               x.DrawBorder(output)
         }

      }),
   )

	// Start trigger goroutines
	go triggers.KeyStrokeTrigger(elchan)

	output.AltScreen()
	output.ClearScreen()

	layout.Reset(output, data)

	for true {
		msg, ok := <-elchan
		if !ok {
			break
		}


      // Rebuild layout after size change
	   w2, h2, _ := term.GetSize(0)
      if width != w2 || height != h2 - 1{
         width = w2
         height = h2 - 1
         layout = layouts.TwoOneHorizontalSplit(height, width)
         layout.Reset(output, data)
      }


		tg.Step(msg, &data)
		output.MoveCursor(height+1, 0)
	}
}
