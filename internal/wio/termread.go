// ======================================================================
// Author: meisto
// Creation Date: Fri 20 Jan 2023 11:53:49 PM CET
// Description: -
// ======================================================================
package wio

import (
	"bufio"
   "log"
	"os"
)

var buffer []rune
var isReadLine bool = false
var isDisplayLine bool = false
var reader *bufio.Reader = bufio.NewReader(os.Stdin)

var readLineOut chan string = make(chan string)


func ReadByte() rune {
   var x = make([]byte, 3)
   numRead, err := os.Stdin.Read(x)
   if err != nil {
      log.Fatal("Unknown error.")
   }
   
   if numRead == 3 {
      log.Print("Haha")
      return ' '
   }

   a := rune(x[0])

   // Hack(ish) code following
   if isReadLine {

      for a != '\r' {

         if isDisplayLine {print(string(a))}

         buffer = append(buffer, rune(a))

         x = make([]byte, 3)
         numRead, err := os.Stdin.Read(x)
         if err != nil {
            log.Fatal("Unknown error.")
         }
         if numRead != 1 {continue}
         a = rune(x[0])

      }

      readLineOut <- string(buffer)

      // Reset buffer
      isReadLine = false
      buffer = []rune{}

      // Recursive call
      return ReadByte()
   }

	return a
}

func ReadLine(display bool) chan string {
   isReadLine = true
   isDisplayLine = display

	return readLineOut
}
