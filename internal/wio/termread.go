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
   "unicode/utf8"
)

var buffer []rune
var isReadLine bool = false
var isDisplayLine bool = false
var maxLength int = 10
var reader *bufio.Reader = bufio.NewReader(os.Stdin)

var readLineOut chan string = make(chan string)


func ReadByte() rune {
   var x = make([]byte, 3)
   numRead, err := os.Stdin.Read(x)
   if err != nil {
      log.Fatal("Unknown error.")
   }

   log.Print(x)
   log.Print(utf8.DecodeRune(x))
   
   if numRead == 3 {
      return ' '
   }

   a := rune(x[0])

   // Hack(ish) code following
   if isReadLine {
      var buffer string
      for a != '\r' {

         if isDisplayLine {print(string(a))}

         buffer += string(a)

         x := make([]byte, 3)
         numRead, err := os.Stdin.Read(x)
         if err != nil {
            log.Fatal("Unknown error.")
         }
         if numRead != 1 { continue }

         a = rune(x[0])
         if utf8.RuneCountInString(buffer) >= maxLength - 1 { 
            buffer += string(a)
            break 
         }

      }

      readLineOut <- buffer

      // Reset buffer
      isReadLine = false

      // Recursive call
      return ReadByte()
   }

	return a
}

func ReadLine(display bool, max int) string {
   isReadLine = true
   isDisplayLine = display
   maxLength = max


	return <- readLineOut
}
