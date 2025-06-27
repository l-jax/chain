package tui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type detailModel struct {
	item     *Item
	viewport *viewport.Model
	quitting bool
	err      error
}

func newDetail() detailModel {
	v := viewport.New(40, 8)
	return detailModel{
		viewport: &v,
	}
}

func (m detailModel) Init() tea.Cmd {
	return nil
}

func (m detailModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case detailMsg:
		m.item = msg.item
		m.viewport.SetContent(msg.item.Text())
	}

	return m, nil
}

func (m detailModel) View() string {
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

func (m detailModel) headerView() string {
	if m.item == nil {
		return "..."
	}
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		headerStyle.Render(m.item.Title()),
		labelStyle.Background(labelColor[m.item.Label()]).Render(m.item.Label()),
	)
}
