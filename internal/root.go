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

   "waelder/internal/config"
   "waelder/internal/io"
   "waelder/internal/language"
   ds "waelder/internal/datastructures"
   "waelder/internal/modes"
   "waelder/internal/layouts"
)

type model struct {
   TextInput      textinput.Model
   Bar            progress.Model

   DatabaseHandle *sql.DB
   Data           ds.Data
   windowHeight   int
   windowWidth    int

   layout         layouts.Layout
   previousMode   modes.Mode
   mode           modes.Mode

   header         string
   footer         string
}


func InitialModel(dbHandle *sql.DB) model {

   ti := textinput.New()
   ti.Blur()
   ti.CharLimit = 10
   ti.Width = 10
   ti.Prompt = ""
   
   root := "/home/tobias/Documents/code/golang/src/waelder"


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

      layout:        layouts.TwoThirdsHorizontalSplit,

      mode:          modes.StartMode,
      header:        "Charakter\n",

      footer: config.StyleDead.Render(strings.Join(language.GetEn().HelpBar, " | ")) +
      "\n" + ti.View(),
   }
}

func (m model) Init() tea.Cmd { return tea.EnterAltScreen }

func (m *model) changeMode(mode modes.Mode) {
   m.previousMode = m.mode
   m.mode = mode
}

func (m *model) switchModes() {
   x := m.mode
   
   m.mode = m.previousMode
   m.previousMode = x
}

func (m model) View() string {

   // Capture StartMode early
   if m.mode == modes.StartMode {return "Welcome"}

   
   // Header
   header := "Charakter\n"

   // Footer
   footer := 
      config.StyleDead.Render(strings.Join(language.GetEn().HelpBar, " | ")) +
      "\n" + m.TextInput.View()

   // Main Content
   var body string
   wh := m.windowHeight - 
      strings.Count(header, "\n") -
      strings.Count(footer, "\n") - 1

   /*
   switch m.mode{
      case modes.LoadingMode :
         body = modes.LoadingView(
            &m,
            wh,
            m.windowWidth,
         )
      case modes.ActiveMode:
         body = modes.ActiveView(
            m.Data,
            wh,
            m.windowWidth,
         )
      case modes.ChoiceMode:
         body = cp.ChoiceView(
            m.Data,
            wh,
            m.windowWidth,
         )
      default:
         return "ERROR 404: page not found ;)\nNeed to add to root page."
   }
   */
   body = m.layout.Display(m.Data)


   /*
   // Cut off body if it is too long
   if strings.Count(body, "\n") > wh - 1 {
      t := strings.Split(body, "\n")
      body = strings.Join(t[len(t) - wh - 1:], "\n")
   }
   */

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

   return m.header + body + padding + m.footer
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
   var cmd tea.Cmd

   switch msg := msg.(type) {

   case tea.WindowSizeMsg:
      m.windowHeight = msg.Height
      m.windowWidth  = msg.Width

      // Only exit StartMode once the first resize message is received.
      if m.mode == modes.StartMode {
         m.layout = m.layout.Resize(m.windowHeight - 3, m.windowWidth)

         m.changeMode(modes.ActiveMode)}
         m.previousMode = modes.ActiveMode // Stop from jumping back to start
   
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
               case modes.LoadingMode:  m.changeMode(modes.LoadingMode)
               case modes.MainMode:     m.changeMode(modes.ActiveMode)
               case modes.ActiveMode:   m.changeMode(modes.MainMode)
            }

         default: // Pass message on to other views

            // Other mode
            switch m.mode {
               case modes.LoadingMode:  // modes.LoadingUpdate(&m.Data, msg)
               case modes.MainMode:     // modes.MainUpdate(&m.Data, msg)
               case modes.ActiveMode:   modes.ActiveUpdate(&m.Data, msg)
               case modes.ChoiceMode:   // cp.ChoiceUpdate(&m.Data, msg)
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

