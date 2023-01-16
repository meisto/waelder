// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 16 Jan 2023 02:02:37 AM CET
// Description: -
// ======================================================================
package model

import(

   tea "github.com/charmbracelet/bubbletea"

   // "dntui/internal/language"
   "dntui/internal/views"
)


// View method
func loadingView(
   model          *model,
   windowHeight   int,
   windowWidth    int,
) string {
   // Data has not yet been loaded
   data := model.Data


   if data.SQLitePath == "" {
      content := "Could not detect environment variable 'dntui_path'."
      return views.PopUp([]string{content}, []string{}, 30, 30)
   }


   content := []string{"Placeholder"}
   // content = append(content, fmt.Sprintf("%t", data.ActiveFiles[0]))
   return views.PopUp(content, []string{}, 30, 30)
}


func loadingUpdate(m *model, msg tea.KeyMsg) {
   /*
   var c0 = cm.Character{Hp: rand.Intn(20), Max_hp: 20, Name: "SomeName01", Initiative: 50}
   var c1 = cm.Character{rand.Intn(20),20,"SomeName02", 50}
   var c2 = cm.Character{rand.Intn(20),20,"SomeName03", 50}
   var c3 = cm.Character{rand.Intn(20),20,"SomeName04", 50}
   var c4 = cm.Character{rand.Intn(20),20,"SomeName05",2}
   var c5 = cm.Character{rand.Intn(20),20,"SomeName06",2}
   var c6 = cm.Character{rand.Intn(20),20,"SomeName07",2}
   var c7 = cm.Character{rand.Intn(20),20,"SomeName08",2}
   var c8 = cm.Character{rand.Intn(20),20,"SomeName09",2}
   var c9 = cm.Character{rand.Intn(20),20,"SomeName10",2}

   data := cm.Data{
      Npcs:    []cm.Character{c1,c2,c0},
      Pcs:     []cm.Character{c3, c9, c8},
      Others:  []cm.Character{c4,c5,c6, c7},
   }
   m.Data = data 
   */
}
