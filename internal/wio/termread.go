// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Fri 20 Jan 2023 11:53:49 PM CET
// Description: -
// ======================================================================
package wio

import (
	"bufio"
	"os"
)

var buffer []rune
var isReadLine bool = false
var isDisplayLine bool = false
var reader *bufio.Reader = bufio.NewReader(os.Stdin)

var readLineOut chan string = make(chan string)


func ReadByte() rune {
	a, _, _ := reader.ReadRune()

   // Hack(ish) code following
   if isReadLine {

      for a != '\r' {

         if isDisplayLine {print(string(a))}

         buffer = append(buffer, rune(a))

         a, _, _ = reader.ReadRune()
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
