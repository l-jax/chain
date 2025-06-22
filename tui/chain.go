package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

type Chain struct {
	rootLink Link
	list     *list.List
	loaded   bool
	quitting bool
	err      error
}

func InitChain(rootLink Link) *Chain {
	l := list.New("branch-1", "branch-2", "branch-3").
		Enumerator(statusEnumerator).
		EnumeratorStyle(enumeratorStyle).
		ItemStyle(itemStyle)

	m := Chain{
		rootLink: rootLink,
		list:     l,
	}

	m.loaded = true
	return &m
}

func statusEnumerator(items list.Items, i int) string {
	return items.At(i).Value()
}

func (m Chain) Init() tea.Cmd {
	return nil
}

func (m Chain) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	return m, cmd
}

func (m Chain) View() string {
	if m.quitting {
		return ""
	}

	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	if !m.loaded {
		return "Loading..."
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render(m.rootLink.Title()),
		bodyStyle.Render(m.rootLink.Description()),
		m.list.String(),
	)
}
