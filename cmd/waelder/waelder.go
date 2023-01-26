// ======================================================================
// Author: Tobias Meisel (meisto)
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

	// open log file
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	defer logFile.Close()

	log.SetOutput(logFile)

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

	defer output.ExitAltScreen()

	ch := make(chan string)

	root.NewRun(dbHandle, ch, output)

	/*
	   p := tea.NewProgram(root.InitialModel(dbHandle))
	   if _, err := p.Run(); err != nil {
	      fmt.Printf("Error: %v", err)
	      os.Exit(1)
	   }
	*/
}
