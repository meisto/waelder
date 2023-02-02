// ======================================================================
// Author: meisto
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
	"waelder/internal/wio"
	"waelder/internal/layouts"
   "waelder/internal/modes"
	"waelder/internal/events"
)

func NewRun(dbHandle *sql.DB, elchan chan string, output *termenv.Output) {
   log.Print("Start Run")

	// TODO: Replace the defaults with a loading procedure
	root := "/home/tobias/Documents/code/golang/src/waelder"
   character1 := wio.LoadCharacterFromFile(path.Join(root, "data/char1.json"))
	// character2 := wio.LoadCharacterFromFile(path.Join(root, "data/char2.json"))
	// character3 := wio.LoadCharacterFromFile(path.Join(root, "data/char3.json"))
	// character4 := wio.LoadCharacterFromFile(path.Join(root, "data/char4.json"))
	character5 := wio.LoadCharacterFromFile(path.Join(root, "data/char5.json"))
	character6 := wio.LoadCharacterFromFile(path.Join(root, "data/char6.json"))
	character7 := wio.LoadCharacterFromFile(path.Join(root, "data/char7.json"))

	f := func(c ds.Character) ds.Character {
		return wio.SyncCharacterWithDatabase(dbHandle, c)
	}

	data := ds.GetData()

   data.AddPlayer(f(character1))
   // data.AddPlayer(f(character2))
   // data.AddPlayer(f(character3))
   // data.AddPlayer(f(character4))

   data.AddAlly(f(character5))
   data.AddEnemy(f(character6))
   data.AddNeutral(f(character7))
   data.PrepareNextRound()

	width, height, _ := term.GetSize(0)
	height -= 1

   // Set initial layout
	layout := layouts.TwoOneHorizontalSplit(height, width)


   InitialNode := data.Graph.Root
   tg := &(data.Graph)

   add := func(n1 ds.Node, key string, n2 ds.Node, g func(), desc string) {
      a := &(data.Graph)
      a.AddEdge( ds.GetEdge(n1, key, n2, g, desc))
   }

   /*
   add(InitialNode, "n", InitialNode, func() {
      data.Step()

      // Redraw fields with this mode
      layout.UpdateMode(output, data, modes.ActiveMode)
   }, "tmp")
   */

   // TODO: Placeholders
   ActionNode := ds.GetNode()
   add(InitialNode, "a", ActionNode, func(){
      layout.UpdateMode(output, data, modes.HelpMode)
   }, "Other Action")

   add(ActionNode, "d", InitialNode, func(){}, "dash")
   add(ActionNode, "f", InitialNode, func(){}, "disengage")
   add(ActionNode, "g", InitialNode, func(){}, "dodge")
   add(ActionNode, "h", InitialNode, func(){}, "help")
   add(ActionNode, "j", InitialNode, func(){}, "hide")
   add(ActionNode, "k", InitialNode, func(){}, "ready")
   add(ActionNode, "l", InitialNode, func(){}, "search")
   add(ActionNode, "v", InitialNode, func(){}, "use object")


   add(InitialNode, "b", InitialNode, func(){}, "React")
   add(InitialNode, "n", InitialNode, func(){}, "Bonus Action")
   add(InitialNode, "m", InitialNode, func(){}, "Apply State")
   add(InitialNode, "y", InitialNode, func(){}, "Redo Action")
   add(InitialNode, "x", InitialNode, func(){}, "Undo Action.")


   // Add popup windows
   addActionPopupSequence(
      output,
      &data,
      InitialNode,
      InitialNode,
      width / 2 - 25, // x
      height / 2, // y
      50, // width
      layout,
   )

   { // Markdown viewer
      n2 := ds.GetNode()
      add(InitialNode, "i", n2, func() {
         var f *layouts.Field = &layout.Fields[1]

         f.SetBorder(layouts.DoubleBorderStyle.Style("darkGreenFg"))
         f.DrawBorder(output)
         layout.UpdateMode(output, data, modes.HelpMode)
      }, "Focus Stat Block")

      add(n2, "k", n2, func() {
         layout.Fields[1].ScrollUp()
         layout.Fields[1].DrawContent(output, data)
         layout.Fields[1].DrawBorder(output)
      }, "Scroll Up")

      add(n2, "j", n2, func() {
         layout.Fields[1].ScrollDown()
         layout.Fields[1].DrawContent(output, data)
         layout.Fields[1].DrawBorder(output)
      }, "Scroll Down")

      add(n2, "i", InitialNode, func() {
         var f *layouts.Field = &layout.Fields[1]

         f.SetBorder(layouts.DoubleBorderStyle)
         f.DrawBorder(output)
         layout.UpdateMode(output, data, modes.HelpMode)
      }, "Relinquish focus")
   }

	// Start trigger goroutines
	go events.KeyStrokeEvent(elchan)


   // Setup screen
	output.AltScreen()
	output.ClearScreen()
	layout.Reset(output, data)

   // Main loop
	for !data.IsCombatOver() {
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

      // Pass the event to the main traversal graph
		tg.Step(msg, &data)

      // Move Cursor to resting position
		output.MoveCursor(height+1, 0)
	}
}
