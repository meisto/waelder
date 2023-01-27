// ======================================================================
// Author: meisto
// Creation Date: Thu 19 Jan 2023 11:42:04 PM CET
// Description: -
// ======================================================================
package main

import (
   "fmt"
   "os"
   
   "golang.org/x/term"
   "azul3d.org/engine/keyboard"

   "github.com/muesli/termenv"
)

func main() {

   output := termenv.NewOutput(os.Stdout)
   output.AltScreen()

   s := output.String("Hello World")

   fmt.Println(s)


   output.ExitAltScreen()
   print(int(os.Stdin.Fd()))


   a, b, _ := term.GetSize(0)
   println(a, " ", b)

   for true {
      keyboard.

   }


}





