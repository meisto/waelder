// ======================================================================
// Author: meisto
// Creation Date: Fri 20 Jan 2023 11:19:11 PM CET
// Description: -
// ======================================================================
package events

import (
   "fmt"

	"waelder/internal/wio"
)

func KeyStrokeEvent(ch chan<- string) {

	triggerMap := make(map[rune]string)

   // These are the possible special inputs
	triggerMap[rune('\t')] = "<TAB>"
	triggerMap[rune('\r')] = "<ENTER>"

   // These are the possible normal inputs
   keys := []byte{
      'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 
      'N', 'P', 'O', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
      'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 
      'n', 'p', 'o', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
      '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
      '+', '-',
   }
   for _, i := range(keys) { triggerMap[rune(i)] = fmt.Sprintf(string(i)) }


	for true {
		b := wio.ReadByte()

		if b == rune('q') {
			close(ch)

			return

		}

		a, exists := triggerMap[b]
		if exists {
			ch <- a
		}
	}
}
