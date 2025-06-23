package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Open struct {
	list     list.Model
	err      error
	loaded   bool
	quitting bool
}

func InitOpen(links []Link) tea.Model {
	m := Open{list: list.New([]list.Item{}, list.NewDefaultDelegate(), windowSize.Width/divisor, windowSize.Height-divisor)}

	m.list.SetShowHelp(false)
	m.list.Title = "pull requests"
	m.list.Styles.Title = titleStyle
	m.list.Styles.NoItems = bodyStyle

	items := make([]list.Item, len(links))
	for i, link := range links {
		items[i] = link
	}
	m.list.SetItems(items)

	m.loaded = true
	return &m
}

func (m Open) Init() tea.Cmd {
	return nil
}

func (m Open) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width/divisor, msg.Height-divisor)
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Open) View() string {
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

func (m Open) SelectedId() uint {
	if item, ok := m.list.SelectedItem().(Link); ok {
		return item.id
	}
	return 0
}
