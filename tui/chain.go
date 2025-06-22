package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Chain struct {
	chain    []Link
	loaded   bool
	quitting bool
	err      error
}

func InitChain(chain []Link) *Chain {
	m := Chain{chain: chain}

	m.loaded = true
	return &m
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
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			titleStyle.Render(m.chain[0].Title()),
			labelStyle.Render(m.chain[0].Label().String()),
		),
		bodyStyle.Render(m.chain[0].Body()),
	)
}
