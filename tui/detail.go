package tui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Detail struct {
	item     *Item
	viewport *viewport.Model
	quitting bool
	err      error
}

func NewDetail() Detail {
	v := viewport.New(40, 8)
	return Detail{
		viewport: &v,
	}
}

func (m Detail) Init() tea.Cmd {
	return nil
}

func (m Detail) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case detailMsg:
		m.item = msg.item
		m.viewport.SetContent(msg.item.Text())
	}

	return m, nil
}

func (m Detail) View() string {
	if m.quitting {
		return ""
	}

	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.headerView(),
		"",
		m.viewport.View(),
	)
}

func (m Detail) headerView() string {
	if m.item == nil {
		return "..."
	}
	return titleStyle.Render(m.item.Title())
}
