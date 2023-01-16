// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 21 Nov 2022 04:14:45 PM CET
// Description: -
// ======================================================================
package model

import (
   "database/sql"
   "strings"

   tea "github.com/charmbracelet/bubbletea"
   "github.com/charmbracelet/bubbles/progress"
   "github.com/charmbracelet/bubbles/textinput"

   // mainWindow "dntui/main_window.go"
   "dntui/internal/data"
   "dntui/internal/config"
   "dntui/internal/language"
)

type Mode int64
const (
   LoadingMode Mode = iota
   MainMode    Mode = iota
   ActiveMode  Mode = iota
)

type model struct {
   TextInput      textinput.Model
   Bar            progress.Model

   DatabaseHandle *sql.DB
   Data           data.Data
   windowHeight   int
   windowWidth    int

   mode           Mode
   cursor         map[string][]int
}


func InitialModel(dbHandle *sql.DB) model {

   ti := textinput.New()
   ti.Blur()
   ti.CharLimit = 10
   ti.Width = 10
   ti.Prompt = ""


   return model {
      TextInput: ti,
      Bar:  progress.New(progress.WithSolidFill("#000000")),
      mode: LoadingMode,

      DatabaseHandle:   dbHandle,
   }
}

func (m model) Init() tea.Cmd {
   return tea.EnterAltScreen
}


func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
   var cmd tea.Cmd

   // d := &m.Data
   // c := &m.Data.Cursor

   switch msg := msg.(type) {

   case tea.WindowSizeMsg:
      m.windowHeight = msg.Height
      m.windowWidth  = msg.Width
   
   case tea.KeyMsg:
      if !m.TextInput.Focused() {
         switch msg.String(){

         case "q", "ctrl+c": // quit
            return m, tea.Quit

         case ":": // enter input
            m.TextInput.Reset()
            m.TextInput.Focus()

         case "tab":
            // Create custom order
            switch m.mode {
               case LoadingMode: m.mode = LoadingMode
               case MainMode: m.mode = ActiveMode
               case ActiveMode: m.mode = MainMode
            }

         default: // Pass message on to other views

            // Other mode
            switch m.mode {
               case LoadingMode: 
                  loadingUpdate(&m, msg)
               case MainMode: mainUpdate(&m, msg)
               case ActiveMode: activeUpdate(&m, msg)
            }
         }
      } else {
         // Parse textfield input
         switch msg.String() {
         case "enter":
            //TODO
            // res := m.TextInput.Value()
            m.TextInput.Reset()
            m.TextInput.Blur()
         }

      }
   }
   m.TextInput, cmd = m.TextInput.Update(msg)

   // Return the updated model to the Bubble Tea runtime for processing.
   return m, cmd
}

func (m model) View() string {
   if m.mode == LoadingMode {
      return loadingView(&m, m.windowHeight, m.windowWidth)
   }
   
   // Header
   header := "Charakter\n\n"

   // Footer
   var footer string
   lang := language.GetEn()
   helpView := config.StyleDead.Render(strings.Join(lang.HelpBar, " | "))


   footer += helpView + "\n" + m.TextInput.View()

   // Main Content
   var body string
   wh := m.windowHeight - strings.Count(header, "\n") - strings.Count(footer, "\n")
   switch m.mode{

   case ActiveMode:
      body = activeView(
         m.Data,
         wh,
         m.windowWidth,
      )
   default:
      body = mainView(
         m.Data,
         wh,
         m.windowWidth,
         m.cursor["mainMode"],
      )
   }


   return header + body + footer

}
