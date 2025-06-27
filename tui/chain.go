package tui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type chainModel struct {
	items    []*pr
	table    table.Model
	spinner  spinner.Model
	quitting bool
	err      error
}

func newChain() chainModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	columns := []table.Column{
		{Title: "it", Width: 5},
		{Title: "branch", Width: 20},
		{Title: "state", Width: 7},
	}

	t := table.New(
		table.WithFocused(false),
		table.WithHeight(5),
		table.WithWidth(40),
		table.WithColumns(columns),
	)

	style := table.DefaultStyles()
	style.Header = tableHeaderStyle

	t.SetStyles(style)

	return chainModel{
		spinner: s,
		table:   t,
	}
}

func (m chainModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m chainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		return m, m.spinner.Tick

	case chainMsg:
		m.items = msg.items
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

func (m chainModel) View() string {
	if m.quitting {
		return ""
	}

	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	if m.items == nil {
		return m.table.View()
	}

	m.SetItems(m.items)
	return m.table.View()
}

func (m *chainModel) SetItems(items []*pr) {
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
