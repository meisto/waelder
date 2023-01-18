// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 21 Nov 2022 04:14:45 PM CET
// Description: -
// ======================================================================
package root

import (
   "database/sql"
   "path"
   "strings"

   tea "github.com/charmbracelet/bubbletea"
   "github.com/charmbracelet/bubbles/progress"
   "github.com/charmbracelet/bubbles/textinput"

   "dntui/internal/config"
   "dntui/internal/io"
   "dntui/internal/language"
   ds "dntui/internal/datastructures"
   cp "dntui/internal/modes/choicepage"
   modi "dntui/internal/modes/modi"
)

type model struct {
   TextInput      textinput.Model
   Bar            progress.Model

   DatabaseHandle *sql.DB
   Data           ds.Data
   windowHeight   int
   windowWidth    int

   previousMode   modi.Mode
   mode           modi.Mode
}


func InitialModel(dbHandle *sql.DB) model {

   ti := textinput.New()
   ti.Blur()
   ti.CharLimit = 10
   ti.Width = 10
   ti.Prompt = ""
   
   root := "/home/tobias/Documents/code/golang/src/dntui"


   // TODO: Replace the defaults with a loading procedure
   character1 := io.LoadCharacterFromFile(path.Join(root, "data/char1.json"))
   character2 := io.LoadCharacterFromFile(path.Join(root, "data/char2.json"))
   character3 := io.LoadCharacterFromFile(path.Join(root, "data/char3.json"))
   character4 := io.LoadCharacterFromFile(path.Join(root, "data/char4.json"))

   f := func(c ds.Character) ds.Character {
      return io.SyncCharacterWithDatabase(dbHandle, c)
   }

   ch1 := f(character1)
   ch2 := f(character2)
   ch3 := f(character3)
   ch4 := f(character4)

   players := []ds.Character{ch1, ch2, ch3, ch4}

   // Return initial data
   return model {
      TextInput: ti,
      Bar:  progress.New(progress.WithSolidFill("#000000")),
      DatabaseHandle:   dbHandle,

      
      Data:  ds.Data {
         Players:    players,
         Allies:     []ds.Character{},
         Enemies:    []ds.Character{},
         Neutrals:   []ds.Character{},
         CombatLog:  ds.CombatLog {
            PreviousRounds: []ds.Round{},
            Current: ds.CreateRound(players),
         },
      },

      mode:          modi.StartMode,
   }
}

func (m model) Init() tea.Cmd { return tea.EnterAltScreen }

func (m *model) changeMode(mode modi.Mode) {
   m.previousMode = m.mode
   m.mode = mode
}

func (m *model) switchModes() {
   x := m.mode
   
   m.mode = m.previousMode
   m.previousMode = x
}

func (m model) View() string {

   // Capture StartMode
   if m.mode == modi.StartMode {return "Welcome"}

   
   // Header
   header := "Charakter\n"

   // Footer
   var footer string
   lang := language.GetEn()
   helpView := config.StyleDead.Render(strings.Join(lang.HelpBar, " | "))
   footer += helpView + "\n" + m.TextInput.View()

   // Main Content
   var body string
   wh := m.windowHeight - strings.Count(header, "\n") - strings.Count(footer, "\n") - 1
   switch m.mode{
      case modi.LoadingMode :
         body = loadingView(
            &m,
            wh,
            m.windowWidth,
         )
      case modi.ActiveMode:
         body = activeView(
            m.Data,
            wh,
            m.windowWidth,
         )
      case modi.ChoiceMode:
         body = cp.ChoiceView(
            m.Data,
            wh,
            m.windowWidth,
         )
      default:
         return "ERROR 404: page not found ;)\nNeed to add to root page."
   }

   // Cut off body if it is too long
   if strings.Count(body, "\n") > wh - 1 {
      t := strings.Split(body, "\n")
      body = strings.Join(t[len(t) - wh - 1:], "\n")
   }

   // Line count
   lc := strings.Count(body, "\n")
   padding := ""
   {
      c := 0
      if wh > lc {
         c = wh -lc
      }
      padding = strings.Repeat("\n", c)
   }

   return header + body + padding + footer
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
   var cmd tea.Cmd

   switch msg := msg.(type) {

   case tea.WindowSizeMsg:
      m.windowHeight = msg.Height
      m.windowWidth  = msg.Width

      // Only exit StartMode once the first resize message is received.
      if m.mode == modi.StartMode {

         m.changeMode(modi.ChoiceMode)}
         m.previousMode = modi.ActiveMode // Stop from jumping back to start
   
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
               case modi.LoadingMode:  m.changeMode(modi.LoadingMode)
               case modi.MainMode:     m.changeMode(modi.ActiveMode)
               case modi.ActiveMode:   m.changeMode(modi.MainMode)
            }

         default: // Pass message on to other views

            // Other mode
            switch m.mode {
               case modi.LoadingMode:  loadingUpdate(&m.Data, msg)
               case modi.MainMode:     mainUpdate(&m.Data, msg)
               case modi.ActiveMode:   activeUpdate(&m.Data, msg)
               case modi.ChoiceMode:   cp.ChoiceUpdate(&m.Data, msg, m.switchModes)
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

