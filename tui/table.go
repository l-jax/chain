package tui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	tableWidth  = 40
	tableHeight = 5
)

type tableModel struct {
	items   []*Item
	table   table.Model
	spinner spinner.Model
	loading bool
	err     error
}

func newTable() tableModel {
	s := spinner.New(spinner.WithSpinner(spinner.Ellipsis))
	s.Style = spinnerStyle

	columns := []table.Column{
		{Title: "id", Width: 5},
		{Title: "branch", Width: 20},
		{Title: "state", Width: 7},
	}

	t := table.New(
		table.WithFocused(false),
		table.WithHeight(tableHeight),
		table.WithWidth(tableWidth),
		table.WithColumns(columns),
	)

	style := table.DefaultStyles()
	style.Header = style.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(purple).
		BorderBottom(true).
		Bold(false)
	style.Selected = style.Selected.
		Foreground(grey).
		Background(purple).
		Bold(false)
	t.SetStyles(style)

	return tableModel{
		spinner: s,
		table:   t,
	}
}

func (m tableModel) Init() tea.Cmd {
	return nil
}

func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case detailMsg:
		m.loading = true
		return m, m.spinner.Tick

	case tableMsg:
		m.items = msg.items
		m.SetItems(m.items)
		m.loading = false
		m.table.Focus()

	case resetMsg:
		m.table.SetCursor(0)
		m.table.Blur()

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

func (m tableModel) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	if m.loading {
		m.table.SetRows([]table.Row{
			[]string{"", m.spinner.View(), ""},
		})
		return m.table.View()
	}

	if m.items == nil {
		m.table.SetRows([]table.Row{})
		return m.table.View()
	}

	return m.table.View()
}

func (m *tableModel) SetItems(items []*Item) {
	rows := make([]table.Row, len(items))
	for i, item := range items {
		rows[i] = []string{
			strconv.FormatUint(uint64(item.Id()), 10),
			item.Title(),
			item.Label(),
		}
	}
	m.table.SetRows(rows)
}
