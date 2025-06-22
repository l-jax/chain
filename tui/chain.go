package tui

import (
	"strconv"

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
		m.prepareChain(),
	)
}

func (m Chain) prepareChain() string {
	if len(m.chain) == 0 {
		return "No chain available"
	}

	var chainView string
	for i, link := range m.chain {
		if i == 0 {
			continue
		}
		chainView = lipgloss.JoinVertical(
			lipgloss.Left,
			chainView,
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				labelStyle.Render("#"+strconv.FormatUint(uint64(link.Id()), 10)),
				titleStyle.Render(link.Title()),
				labelStyle.Render(link.Label().String()),
			),
		)
	}
	return chainView
}
