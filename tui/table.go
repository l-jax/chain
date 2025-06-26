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
	return Table{
		spinner: s,
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

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
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

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m.table = t
}
