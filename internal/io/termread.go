// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Fri 20 Jan 2023 11:53:49 PM CET
// Description: -
// ======================================================================
package io

import (
   "bufio"
   "os"
)



func ReadByte() rune {
   reader := bufio.NewReader(os.Stdin)

   a, _ := reader.ReadByte()

   return rune(a)
}  

func ReadLine() string {
   reader := bufio.NewReader(os.Stdin)

   a, _ := reader.ReadString('\n')

   return a
}
