// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 16 Jan 2023 02:02:37 AM CET
// Description: -
// ======================================================================
package modes

import(

   // tea "github.com/charmbracelet/bubbletea"

   // "waelder/internal/language"
   "waelder/internal/views"
)


// View method
func LoadingView(
   windowHeight   int,
   windowWidth    int,
) string {


   content := [][]string{[]string{"Hallo", "dies", "ist", "ein", "test"}, []string{"Dies", "Ist", "Ein", "weiterer", "Test", "adassdfa"}}
   ch := [][]views.ColorHelper{}
   return views.PopUp(content, ch, views.DefaultBorder,0,0)
}


// func LoadingUpdate(m *model, msg tea.KeyMsg) { }
