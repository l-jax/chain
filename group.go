package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type group struct {
	focus  bool
	state  state
	list   list.Model
	height int
	width  int
}

func (g *group) Focus() {
	g.focus = true
}

func (g *group) Blur() {
	g.focus = false
}

func (g *group) Focused() bool {
	return g.focus
}

func newGroup(state state) group {
	var focus bool
	if state == StateOpen {
		focus = true
	}
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	defaultList.SetShowHelp(false)
	return group{focus: focus, state: state, list: defaultList}
}

func (g group) Init() tea.Cmd {
	return nil
}

func (g group) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return g, nil
}

func (g group) View() string {
	return g.getStyle().Render(g.list.View())
}

func (g *group) getStyle() lipgloss.Style {
	if g.Focused() {
		return lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			Height(g.height).
			Width(g.width)
	}
	return lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.HiddenBorder()).
		Height(g.height).
		Width(g.width)
}
