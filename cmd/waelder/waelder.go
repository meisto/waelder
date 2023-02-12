// ======================================================================
// Author: meisto
// Creation Date: Wed 16 Nov 2022 11:18:35 PM CET
// Description: -
// ======================================================================
package main

import (
	"log"
	"os"

	"github.com/muesli/termenv"
	"golang.org/x/term"

	root "waelder/internal"
	"waelder/internal/config"
	"waelder/internal/wio"
)

func main() {
	// Set up logging
	logFilePath := "/tmp/logFile.log"

   { // Setup logging
      logFile, err := os.OpenFile(
         logFilePath, 
         os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644,
      )
      if err != nil {
         log.Panic(err)
      }
      log.SetFlags(log.Lshortfile | log.LstdFlags)
      log.SetOutput(logFile)
      defer logFile.Close()
   }


	// Load database
	dbHandle := wio.GetDatabaseHandle()
	defer dbHandle.Close()

	// Create tables if not existing
	wio.CreateTables(dbHandle)

	// Setup styles
   colorscheme := wio.ReadLocalFileToString("data/settings/color.txt")
	config.SetupStylemap(colorscheme)

	// Switch terminal from cooked to raw
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	output := termenv.NewOutput(os.Stdout)
	output.AltScreen()
	output.ClearScreen()
	defer output.ExitAltScreen()


   data := root.PreCombat(dbHandle, output)
   return
	root.CombatLoop(&data, output)
}
