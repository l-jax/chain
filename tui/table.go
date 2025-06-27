package tui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Table struct {
	items    []*Item
	table    table.Model
	spinner  spinner.Model
	quitting bool
	err      error
}

func NewTable() Table {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	t := table.New(
		table.WithFocused(false),
		table.WithHeight(5),
	)

	style := table.DefaultStyles()
	style.Selected = tableSelectedStyle
	style.Header = tableHeaderStyle

	t.SetStyles(style)

	return Table{
		spinner: s,
		table:   t,
	}
}

func (m Table) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m Table) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		return m, m.spinner.Tick

	case detailMsg:
		m.items = msg.linked
		m.table.Focus()

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	}
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

func (m Table) View() string {
	if m.quitting {
		return ""
	}

	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	if m.items == nil {
		return m.spinner.View()
	}

	m.SetItems(m.items)

	return m.table.View()
}

func (m *Table) SetItems(items []*Item) {
	columns := []table.Column{
		{Title: "ID", Width: 5},
		{Title: "Branch", Width: 20},
		{Title: "State", Width: 10},
	}
	rows := make([]table.Row, len(items))
	for i, item := range items {
		rows[i] = []string{
			strconv.FormatUint(uint64(item.Id()), 10),
			item.Title(),
			item.Label(),
		}
	}
	m.table.SetColumns(columns)
	m.table.SetRows(rows)
}
