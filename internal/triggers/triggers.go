// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Fri 20 Jan 2023 11:19:11 PM CET
// Description: -
// ======================================================================
package triggers

import (
   "waelder/internal/io"
)


func KeyStrokeTrigger(ch chan<- string) {

   triggerMap := make(map[rune]string)
   triggerMap[rune('n')] = "n"
   triggerMap[rune('\t')] = "<TAB>"

   for true {
      b := io.ReadByte()

      if b == rune('q') {
         close(ch)

         return

      }

      a, exists := triggerMap[b]
      if exists { ch <- a }
   }
}



