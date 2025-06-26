package tui

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type Detail struct {
	item     *Item
	linked   []*Item
	quitting bool
	err      error
}

func (m Detail) Init() tea.Cmd {
	return nil
}

func (m Detail) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case detailMsg:
		m.item = msg.item
		m.linked = msg.linked
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

	if m.item == nil {
		return "Loading..."
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			titleStyle.Render("chain"),
			bodyStyle.Padding(0, 2).Render(m.item.Title()),
			labelStyle.Render(m.item.Label()),
		),
		bodyStyle.Padding(1, 0).Render(m.item.Text()),
		m.RenderChain(),
	)
}

func (m *Detail) RenderChain() string {
	rows := make([][]string, len(m.linked))
	for i, item := range m.linked {
		rows[i] = []string{
			strconv.FormatUint(uint64(item.Id()), 10),
			item.Title(),
			item.Label(),
		}
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(purple)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			case row%2 == 0:
				return evenRowStyle
			default:
				return oddRowStyle
			}
		}).
		Headers("id", "branch", "state").
		Rows(rows...)

	return t.Render()
}
