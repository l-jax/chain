package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type List struct {
	list     list.Model
	err      error
	loaded   bool
	quitting bool
}

func InitList(items []*Item) tea.Model {
	m := List{list: list.New([]list.Item{}, list.NewDefaultDelegate(), windowSize.Width/divisor, windowSize.Height-divisor)}

	m.list.SetShowHelp(false)
	m.list.Title = "pull requests"
	m.list.Styles.Title = titleStyle
	m.list.Styles.NoItems = bodyStyle

	listItems := make([]list.Item, len(items))
	for i, pr := range items {
		listItems[i] = pr
	}
	m.list.SetItems(listItems)

	m.loaded = true
	return &m
}

func (m List) Init() tea.Cmd {
	return nil
}

func (m List) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width/divisor, msg.Height-divisor)
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m List) View() string {
	if m.quitting {
		return ""
	}

	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	if !m.loaded {
		return "Loading..."
	}

	return m.list.View()
}

func (m List) SelectedId() uint {
	if item, ok := m.list.SelectedItem().(*Item); ok {
		return item.Id()
	}
	return 0
}
