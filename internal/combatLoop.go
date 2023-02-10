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

func NewRun(dbHandle *sql.DB, eventChannel chan string, output *termenv.Output) {
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
	output.AltScreen()
	output.ClearScreen()
	layout := layouts.TwoOneHorizontalSplit(height, width, output, data)

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


   { // Add actions
      ActionNode := ds.GetNode()
      add(InitialNode, "a", ActionNode, func(){
         layout.UpdateMode(data, modes.HelpMode)
      }, "Action")

      f := func(key string, g func(), desc string) {
         add(ActionNode, key, InitialNode, g, desc)
      }
      f( "a", func() {getAttack(&data, layout, eventChannel)}, "Attack")
      f( "d", func() {getDash(&data,layout)}, "dash")
      f( "f", func() {getDisengage(&data, layout)}, "disengage")
      f( "g", func() {getDodge(&data, layout)}, "dodge")
      f( "h", func() {getHelp(&data, layout)}, "help")
      f( "j", func() {getHide(&data, layout)}, "hide")
      f( "k", func() {getReady(&data, layout)}, "ready")
      f( "l", func() {getSearch(&data, layout)}, "search")
      f( "v", func() {getUseObject(&data, layout)}, "use object")
   }


   placeholder := func() {data.StepWoAction(); layout.Reset(data)}
   add(InitialNode, "b", InitialNode, placeholder, "React")
   add(InitialNode, "n", InitialNode, placeholder, "Bonus Action")
   add(InitialNode, "m", InitialNode, placeholder, "Apply State")
   add(InitialNode, "y", InitialNode, placeholder, "Redo Action")
   add(InitialNode, "x", InitialNode, placeholder, "Undo Action.")


   { // Markdown viewer
      n2 := ds.GetNode()
      add(InitialNode, "i", n2, func() {
         var f *layouts.Field = &layout.Fields[1]

         f.SetBorder(layouts.DoubleBorderStyle.Style("darkGreenFg"))
         f.DrawBorder()
         layout.UpdateMode(data, modes.HelpMode)
      }, "Focus Stat Block")

      add(n2, "k", n2, func() {
         layout.Fields[1].ScrollUp()
         layout.Fields[1].UpdateContent(data)
         layout.Fields[1].DrawContent()
         layout.Fields[1].DrawBorder()
      }, "Scroll Up")

      add(n2, "j", n2, func() {
         layout.Fields[1].ScrollDown()
         layout.Fields[1].UpdateContent(data)
         layout.Fields[1].DrawContent()
         layout.Fields[1].DrawBorder()
      }, "Scroll Down")

      add(n2, "i", InitialNode, func() {
         var f *layouts.Field = &layout.Fields[1]

         f.SetBorder(layouts.DoubleBorderStyle)
         f.DrawBorder()
         layout.UpdateMode(data, modes.HelpMode)
      }, "Relinquish focus")
   }

	// Start trigger goroutines
	go events.KeyStrokeEvent(eventChannel)

   // Main loop
	for !data.IsCombatOver() {
      layout.Reset(data)
		msg, ok := <- eventChannel
		if !ok {
			break
		}

      // Rebuild layout after size change
      /*
	   w2, h2, _ := term.GetSize(0)
      if width != w2 || height != h2 - 1{
         width = w2
         height = h2 - 1
         layout = layouts.TwoOneHorizontalSplit(height, width, output)
         layout.Reset(data)
      }
      */

      // Pass the event to the main traversal graph
		tg.Step(msg, &data)

      // Move Cursor to resting position
		output.MoveCursor(height+1, 0)
	}
}
