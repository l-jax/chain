package tui

import (
	"sort"
	"strconv"

	"github.com/charmbracelet/bubbles/key"
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
	loading bool
	err     error
}

func newTable() tableModel {
	columns := []table.Column{
		{Title: "id", Width: 5},
		{Title: "branch", Width: 20},
		{Title: "state", Width: 7},
		{Title: " ", Width: 1},
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
		Foreground(lightGrey).
		Background(purple).
		Bold(false)
	t.SetStyles(style)

	return tableModel{
		table: t,
	}
}

func (m tableModel) Init() tea.Cmd {
	return nil
}

func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Up):
			m.table, _ = m.table.Update(msg)
			if m.table.Focused() {
				cmds = append(cmds, m.focusItem)
			}
		case key.Matches(msg, keys.Down):
			m.table, _ = m.table.Update(msg)
			if m.table.Focused() {
				cmds = append(cmds, m.focusItem)
			}
		}

	case tableLoadMsg:
		m.loading = true

	case tableMsg:
		m.items = msg.items
		m.SetItems(m.items)
		m.loading = false
		m.table.Focus()
		m.table.SetCursor(0)

	case resetMsg:
		m.table.Blur()
	}

	return m, tea.Batch(cmds...)
}

func (m tableModel) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	if m.loading {
		m.table.SetRows([]table.Row{
			{"?", "?", "?", "?"},
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
	sort.Slice(items, func(i, j int) bool {
		return items[i].DependsOn() > items[j].DependsOn()
	})

	rows := make([]table.Row, len(items))
	for i, item := range items {
		var ok string
		if item.Blocked() {
			ok = "x"
		} else if !item.HasTargetLabel() {
			ok = "-"
		} else {
			ok = "✔"
		}

		rows[i] = []string{
			strconv.FormatUint(uint64(item.Id()), 10),
			item.Description(),
			item.State(),
			ok,
		}
	}
	m.table.SetRows(rows)
}

func (m *tableModel) focusItem() tea.Msg {
	return detailMsg{
		item: m.items[m.table.Cursor()],
	}
}
