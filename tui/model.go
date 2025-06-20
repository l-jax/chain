package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const divisor = 4

type Model struct {
	lists    []list.Model
	focussed index
	loaded   bool
	err      error
	quitting bool
}

func NewModel() *Model {
	return &Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if m.loaded {
			break
		}
		m.InitLists(msg.Width, msg.Height)
		listStyle.Width(msg.Width / divisor)
		focussedStyle.Width(msg.Width / divisor)
		listStyle.Height(msg.Height - divisor)
		focussedStyle.Height(msg.Height - divisor)

		m.loaded = true
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "left", "h":
			m.SetFocussed(active)
		case "right", "l":
			m.SetFocussed(chain)
		}
	}

	var cmd tea.Cmd
	m.lists[m.focussed], cmd = m.lists[m.focussed].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}

	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	if !m.loaded {
		return "Loading..."
	}

	activeView := m.lists[active].View()
	chainView := m.lists[chain].View()

	switch m.focussed {
	default:
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			focussedStyle.Render(activeView),
			listStyle.Render(chainView),
		)
	case chain:
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			listStyle.Render(activeView),
			focussedStyle.Render(chainView),
		)
	}
}

func (m *Model) InitLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divisor, height/2)
	defaultList.SetShowHelp(false)

	m.lists = []list.Model{defaultList, defaultList}

	m.lists[active].Title = "ACTIVE"
	m.lists[chain].Title = "CHAIN"

	m.lists[active].SetItems(
		[]list.Item{
			NewLink("Active Link 1", "Description for active link 1", active),
			NewLink("Active Link 2", "Description for active link 2", active),
		},
	)

	m.lists[chain].SetItems(
		[]list.Item{
			NewLink("Chain Link 1", "Description for chain link 1", chain),
			NewLink("Chain Link 2", "Description for chain link 2", chain),
		},
	)
}

func (m *Model) SetFocussed(idx index) {
	if idx >= index(len(m.lists)) {
		return
	}
	m.focussed = idx
}
