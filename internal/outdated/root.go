// ======================================================================
// Author: meisto
// Creation Date: Mon 21 Nov 2022 04:14:45 PM CET
// Description: -
// ======================================================================
package root

import (
	"database/sql"
	"path"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"waelder/internal/config"
	ds "waelder/internal/datastructures"
	"waelder/internal/graph"
	"waelder/internal/io"
	"waelder/internal/language"
	"waelder/internal/layouts"
	"waelder/internal/modes"
)

type model struct {
	TextInput textinput.Model

	DatabaseHandle *sql.DB
	Data           ds.Data
	windowHeight   int
	windowWidth    int

	layout       layouts.Layout
	previousMode modes.Mode
	mode         modes.Mode

	header string
	footer string

	traversalGraph graph.Graph
}

func InitialModel(dbHandle *sql.DB) model {

	// TODO: Replace the defaults with a loading procedure
	root := "/home/tobias/Documents/code/golang/src/waelder"
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

	var InitialNode graph.Node = graph.GetNode()

	// Window resize related nodes
	var ResizeNode graph.Node = graph.GetNode()

	// RootNodes
	var RedrawNode graph.Node = graph.GetNode()

	// Layout specific nodes
	var LayoutNode graph.Node = graph.GetNode()
	var Layout1Node graph.Node = graph.GetNode()
	var Layout2Node graph.Node = graph.GetNode()
	var Layout3Node graph.Node = graph.GetNode()

	var allNodes []graph.Node = []graph.Node{InitialNode, ResizeNode, RedrawNode, LayoutNode, Layout1Node, Layout2Node, Layout3Node}

	tg := graph.GetGraph(InitialNode)
	for _, i := range allNodes {
		tg.AddNode(i)
	}

	ti := textinput.New()
	ti.Blur()
	ti.CharLimit = 10
	ti.Width = 10
	ti.Prompt = ""

	// Return initial data
	return model{
		TextInput:      ti,
		DatabaseHandle: dbHandle,

		Data: ds.Data{
			Players:  players,
			Allies:   []ds.Character{},
			Enemies:  []ds.Character{},
			Neutrals: []ds.Character{},
			CombatLog: ds.CombatLog{
				PreviousRounds: []ds.Round{},
				Current:        ds.CreateRound(players),
			},
		},

		layout: layouts.TwoOneHorizontalSplit,

		mode:   modes.StartMode,
		header: "Charakter\n",

		footer: config.GetStyle("dead").Render(strings.Join(language.GetEn().HelpBar, " | ")) +
			"\n" + ti.View(),

		traversalGraph: tg,
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
	if m.mode == modes.StartMode {
		return "Welcome"
	}

	// Header
	header := "Charakter\n"

	// Footer
	footer :=
		config.GetStyle("dead").Render(strings.Join(language.GetEn().HelpBar, " | ")) +
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
			c = wh - lc
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
		m.windowWidth = msg.Width

		print("[NOTE] Resized the screen to ", m.windowHeight, "h x ", m.windowWidth, "w.")

		// Only exit StartMode once the first resize message is received.
		if m.mode == modes.StartMode {
			m.layout = m.layout.Resize(m.windowHeight-3, m.windowWidth)

			m.changeMode(modes.ActiveMode)
		}
		m.previousMode = modes.ActiveMode // Stop from jumping back to start

	case tea.KeyMsg:
		if !m.TextInput.Focused() {
			switch msg.String() {

			case "q", "ctrl+c": // quit
				return m, tea.Quit

			case ":": // enter input
				m.TextInput.Reset()
				m.TextInput.Focus()

			case "tab":
				// Create custom order
				switch m.mode {
				case modes.LoadingMode:
					m.changeMode(modes.LoadingMode)
				case modes.MainMode:
					m.changeMode(modes.ActiveMode)
				case modes.ActiveMode:
					m.changeMode(modes.MainMode)
				}

			default: // Pass message on to other views

				// Other mode
				switch m.mode {
				case modes.LoadingMode: // modes.LoadingUpdate(&m.Data, msg)
				case modes.MainMode: // modes.MainUpdate(&m.Data, msg)
				case modes.ActiveMode:
					modes.ActiveUpdate(&m.Data, msg)
				case modes.ChoiceMode: // cp.ChoiceUpdate(&m.Data, msg)
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
