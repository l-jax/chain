package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	active list.Model
	chain  *tree.Tree
}

func newModel() model {
	prs := getActivePullRequests()

	t := getTree("some-example")

	m := model{
		active: list.New(prs, list.NewDefaultDelegate(), 0, 0),
		chain:  t,
	}

	m.active.Title = "ACTIVE"
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.active.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.active, cmd = m.active.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top,
		docStyle.Render(m.active.View()),
		docStyle.Render(m.chain.String()))
}
