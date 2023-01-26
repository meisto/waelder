// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Fri 20 Jan 2023 11:19:11 PM CET
// Description: -
// ======================================================================
package triggers

import (
   "fmt"

	"waelder/internal/wio"
)

func KeyStrokeTrigger(ch chan<- string) {

	triggerMap := make(map[rune]string)
	triggerMap[rune('n')] = "n"
	triggerMap[rune('\t')] = "<TAB>"
	triggerMap[rune('p')] = "p"
	triggerMap[rune('o')] = "o"
	triggerMap[rune('\r')] = "<ENTER>"
   keys := []byte{
      'n', 'p', 'o', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
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
