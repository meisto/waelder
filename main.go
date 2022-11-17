// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Wed 16 Nov 2022 11:18:35 PM CET
// Description: -
// ======================================================================

package main

import (
   "fmt"
   "os"
   "math/rand"
   "strings"

   tea "github.com/charmbracelet/bubbletea"
   "github.com/charmbracelet/lipgloss"
   "github.com/charmbracelet/bubbles/progress"
   "github.com/charmbracelet/bubbles/key"
   "github.com/charmbracelet/bubbles/help"
   // "github.com/charmbracelet/bubbles/table"
)

type character struct{
   hp       int
   max_hp   int
   name     string
}

type cursor struct{
   group    int
   element  int
}

type keyMap struct {
   Up    key.Binding
   Down  key.Binding
   Help  key.Binding
   Quit  key.Binding
}

type model struct {
   keys     keyMap
   help     help.Model

   npcs     []character
   pcs      []character
   others   []character
   cursor   cursor
}

const (
   healthcolor100 string   = "#00FF00"
   healthcolor75 string    = "#F3F034"
   healthcolor50 string    = "#DE7007"
   healthcolor25 string    = "#FF0000"

)
/////////////////////////////////////////
func (k keyMap) ShortHelp() []key.Binding {
   return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
   return [][]key.Binding{
      {k.Up, k.Down}, // first column
      {k.Help, k.Quit},                // second column
   }
}

var keys = keyMap{
   Up: key.NewBinding(
      key.WithKeys("up", "k"),
      key.WithHelp("↑/k", "move up"),
   ),
   Down: key.NewBinding(
      key.WithKeys("down", "j"),
      key.WithHelp("↓/j", "move down"),
   ),
   Help: key.NewBinding(
      key.WithKeys("?", "h"),
      key.WithHelp("?/h", "toggle help"),
   ),
   Quit: key.NewBinding(
      key.WithKeys("q", "esc", "ctrl+c"),
      key.WithHelp("q", "quit"),
   ),
}

/////////////////////////////////////////

func initialModel() model {
   var c0 = character{rand.Intn(20),20,"SomeName01"}
   var c1 = character{rand.Intn(20),20,"SomeName02"}
   var c2 = character{rand.Intn(20),20,"SomeName03"}
   var c3 = character{rand.Intn(20),20,"SomeName04"}
   var c4 = character{rand.Intn(20),20,"SomeName05"}
   var c5 = character{rand.Intn(20),20,"SomeName06"}
   var c6 = character{rand.Intn(20),20,"SomeName07"}
   var c7 = character{rand.Intn(20),20,"SomeName08"}
   var c8 = character{rand.Intn(20),20,"SomeName09"}
   var c9 = character{rand.Intn(20),20,"SomeName10"}

   return model{
      keys:    keys,
      help:    help.New(),
      npcs:    []character{c1,c2,c0},
      pcs:     []character{c3, c9, c8},
      others:  []character{c4,c5,c6, c7},
      cursor:  cursor{0,0},
   }
}

func (m model) Init() tea.Cmd {
   return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
   switch msg := msg.(type) {

   case tea.KeyMsg:
      switch {
      case key.Matches(msg, m.keys.Quit):
         return m, tea.Quit
      case key.Matches(msg, m.keys.Up):
         if m.cursor.element > 0 {
            m.cursor.element--
         } else {
            // Update group
            m.cursor.group = (m.cursor.group - 1) % 3
            if m.cursor.group < 0 {m.cursor.group += 3}

            // Update element
            l := [][]character{m.npcs, m.pcs, m.others}
            m.cursor.element = len(l[m.cursor.group]) - 1
            if m.cursor.element < 0 {m.cursor.element = 0}
         }

         // The "down" and "j" keys move the cursor down
      case key.Matches(msg, m.keys.Down):
         l := [][]character{m.npcs, m.pcs, m.others}

         if m.cursor.element < len(l[m.cursor.group])-1 {
            m.cursor.element++
         } else {
            m.cursor.group = (m.cursor.group + 1) % 3
            m.cursor.element = 0
         }

         // The "enter" key and the spacebar (a literal space) toggle
         // the selected state for the item that the cursor is pointing at.
      case key.Matches(msg, m.keys.Help):
         m.help.ShowAll = !m.help.ShowAll
      }
   }

      // Return the updated model to the Bubble Tea runtime for processing.
      // Note that we're not returning a command.
      return m, nil
   }

   var (
      style1      = lipgloss.NewStyle().Background(lipgloss.Color("#84AF87")).Foreground(lipgloss.Color("#000000"))
      styleDead   = lipgloss.NewStyle().Foreground(lipgloss.Color("#444444"))
      bar = progress.New(progress.WithSolidFill("#000000"))
   )


   // View method
   func (m model) View() string {
      // Bar setup
      bar.ShowPercentage = false
      bar.Width = 10


      // The header
      s := "Charakter\n\n"

      // Iterate over our choices
      for i1, x := range [3][]character{m.npcs, m.pcs, m.others} {

         heading := "DEFAULT HEADING"
         switch i1 {
            case 0: heading = "NPCS"
            case 1: heading = "PCS"
            case 2: heading = "OTHERS"
         }
         s += style1.Render(fmt.Sprintf(" %s ", heading))
         s += "\n"

         for i2 := 0; i2 < len(x); i2++ {

            char := x[i2]

            cursor := " " // no cursor
            if i1 == m.cursor.group && m.cursor.element == i2 {
               cursor = ">"
            }
            charName := fmt.Sprintf(" %10s ",char.name) // not selected
            comp1 := cursor + charName

            // Generate 
            var hpPercentage float64 = float64(char.hp) / float64(x[i2].max_hp)

            switch {
            case hpPercentage >= 0.75:
               bar.FullColor = healthcolor100
            case hpPercentage >= 0.5:
               bar.FullColor = healthcolor75
            case hpPercentage >= 0.25:
               bar.FullColor = healthcolor50
            default:
               bar.FullColor = healthcolor25
            }

            comp2 := fmt.Sprintf("%s", bar.ViewAs(hpPercentage))
            comp3 := fmt.Sprintf(" %03d/%03d", x[i2].hp, x[i2].max_hp)

            comp4 := " | AC: 1011 | SPEED: 10"

            // Grey out dead characters
            if hpPercentage == 0 {
               comp1 = styleDead.Render(comp1)
               comp2 = styleDead.Render(comp2)
               comp3 = styleDead.Render(comp3)
               comp4 = styleDead.Render(comp4)
            }

            s += comp1 + comp2 + comp3 + comp4 + "\n"
         }

         s += "\n"

      }
      helpView := m.help.View(m.keys)
      height := strings.Count(s, "\n") - strings.Count(helpView, "\n") - 15

      return "\n" + s + strings.Repeat("\n", height) + helpView
   }

   func main() {
      p := tea.NewProgram(initialModel())
      if _, err := p.Run(); err != nil {
         fmt.Printf("Error: %v", err)
         os.Exit(1)
      }
   }
