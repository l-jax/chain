package tui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type chainModel struct {
	items   []*Item
	table   table.Model
	spinner spinner.Model
	loading bool
	err     error
}

func newChain() chainModel {
	s := spinner.New()
	s.Spinner = spinner.Ellipsis
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	columns := []table.Column{
		{Title: "id", Width: 5},
		{Title: "branch", Width: 20},
		{Title: "state", Width: 7},
	}

	t := table.New(
		table.WithFocused(false),
		table.WithHeight(6),
		table.WithWidth(40),
		table.WithColumns(columns),
	)

	return chainModel{
		spinner: s,
		table:   t,
	}
}

func (m chainModel) Init() tea.Cmd {
	return nil
}

func (m chainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case detailMsg:
		m.loading = true
		return m, m.spinner.Tick

	case chainMsg:
		m.items = msg.items
		m.SetItems(m.items)
		m.loading = false

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

func (m chainModel) View() string {
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

func (m *chainModel) SetItems(items []*Item) {
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
