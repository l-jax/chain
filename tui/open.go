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

func InitOpen() tea.Model {
	m := Open{list: list.New([]list.Item{}, list.NewDefaultDelegate(), windowSize.Width/divisor, windowSize.Height/divisor)}

	m.list.SetShowHelp(false)
	m.list.Title = "open"
	m.list.Styles.Title = titleStyle
	m.list.Styles.NoItems = bodyStyle

	m.list.SetItems([]list.Item{
		NewLink("Chain 1", "Description for Chain 1", "branch-1", 1, 1, label(open)),
		NewLink("Chain 2", "Description for Chain 2", "branch-1", 2, 2, label(blocked)),
		NewLink("Chain 3", "Description for Chain 3", "branch-1", 3, 3, label(merged)),
	})

	m.loaded = true
	return &m
}

func (m Open) Init() tea.Cmd {
	return nil
}

func (m Open) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width/divisor, msg.Height/divisor)
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
