// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Fri 10 Feb 2023 10:53:06 PM CET
// Description: -
// ======================================================================
package root

import(
   "time"
   "path"
   "database/sql"

   "github.com/muesli/termenv"

	"golang.org/x/term"

   ds "waelder/internal/datastructures"
	"waelder/internal/layouts"
	"waelder/internal/wio"
)

func PreCombat(dbHandle *sql.DB, output *termenv.Output) ds.Data {
   

	width, height, _ := term.GetSize(0)
	height -= 1

	data := ds.GetData()
   l := layouts.Fullscreen(height, width, output, data)
   output.ClearScreen()
   l.Reset(data)
   time.Sleep(10 * time.Second)

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


   data.AddPlayer(f(character1))
   // data.AddPlayer(f(character2))
   // data.AddPlayer(f(character3))
   // data.AddPlayer(f(character4))

   data.AddAlly(f(character5))
   data.AddEnemy(f(character6))
   data.AddNeutral(f(character7))
   data.PrepareNextRound()



   return data
}
