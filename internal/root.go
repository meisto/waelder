// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Fri 20 Jan 2023 02:00:59 AM CET
// Description: -
// ======================================================================
package root 

import ( 
   "database/sql"
   "path"

   "github.com/muesli/termenv"

   "golang.org/x/term"

   "waelder/internal/io"
   ds "waelder/internal/datastructures"
   "waelder/internal/layouts"
   "waelder/internal/datastructures/graph"
   "waelder/internal/triggers"
   "waelder/internal/modes"
)

func NewRun(dbHandle *sql.DB, elchan chan string, output *termenv.Output) {
   
   // TODO: Replace the defaults with a loading procedure
   root := "/home/tobias/Documents/code/golang/src/waelder"
   character1 := io.LoadCharacterFromFile(path.Join(root, "data/char1.json"))
   character2 := io.LoadCharacterFromFile(path.Join(root, "data/char2.json"))
   character3 := io.LoadCharacterFromFile(path.Join(root, "data/char3.json"))
   character4 := io.LoadCharacterFromFile(path.Join(root, "data/char4.json"))

   f := func(c ds.Character) ds.Character {
      return io.SyncCharacterWithDatabase(dbHandle, c)
   }

   ch1 := f(character1)
   ch2 := f(character2)
   ch3 := f(character3)
   ch4 := f(character4)

   players := []ds.Character{ch1, ch2, ch3, ch4}

   data := ds.Data {
      Players:    players,
      Allies:     []ds.Character{},
      Enemies:    []ds.Character{},
      Neutrals:   []ds.Character{},
      CombatLog:  ds.CombatLog {
         PreviousRounds: []ds.Round{},
         Current: ds.CreateRound(players),
      },
   }


   width, height, _ := term.GetSize(0)
   height -= 1
   layout := layouts.TwoOneHorizontalSplit(height, width)

   // 
   var InitialNode graph.Node = graph.GetNode()


   tg := graph.GetGraph(InitialNode)
   tg.AddEdge(
      graph.GetEdge(
         InitialNode,
         "n",
         InitialNode,
         func() {
            if data.CombatLog.Current.IsDone() {
               // Trigger start of new Round
               oldRound, newRound := data.CombatLog.Current.GetNextRound()

               data.CombatLog.PreviousRounds   = append(data.CombatLog.PreviousRounds, oldRound)
               data.CombatLog.Current          = newRound
            } else {
               data.CombatLog.Current.Step()
            }

            // Redraw fields with this mode
            layout.UpdateMode(output, data, modes.ActiveMode)
         },

      ),
   )






   // Start trigger goroutines
   go triggers.KeyStrokeTrigger(elchan)


   output.AltScreen()
   output.ClearScreen()
   

   layout.DrawBorders(output)
   layout.Display(output, data)
   output.MoveCursor(height + 1, 0)


   for true {
      msg, ok := <- elchan
      if !ok { break }

      tg.Step(msg, &data)
      layout.Display(output, data)
      output.MoveCursor(height + 1, 0)
   }
}
