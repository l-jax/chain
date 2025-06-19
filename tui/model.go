package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const divisor = 5

type Model struct {
	lists    []list.Model
	focussed index
	loaded   bool
	err      error
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
		m.loaded = true
	}
	var cmd tea.Cmd
	m.lists[m.focussed], cmd = m.lists[m.focussed].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	if !m.loaded {
		return "Loading..."
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, m.lists[active].View(), m.lists[chain].View())
}

func (m *Model) InitLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divisor-2, height)
	defaultList.SetShowHelp(false)

	m.lists = []list.Model{defaultList, defaultList}

	m.lists[active].Title = "ACTIVE"
	m.lists[chain].Title = "CHAIN"
}
