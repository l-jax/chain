package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type listModel struct {
	list     list.Model
	err      error
	loaded   bool
	focussed bool
	quitting bool
}

func newList() tea.Model {
	d := list.NewDefaultDelegate()
	d.Styles.NormalTitle = listTitleStyle
	d.Styles.SelectedTitle = selectedTitleStyle
	d.Styles.NormalDesc = listDescStyle
	d.Styles.SelectedDesc = selectedDescStyle

	m := listModel{list: list.New([]list.Item{}, d, 18, 20)}

	m.list.SetShowHelp(false)
	m.list.Title = "pull requests"
	m.list.Styles.Title = listHeaderStyle
	m.list.Styles.NoItems = helpStyle

	m.loaded = true
	return &m
}

func (m listModel) Init() tea.Cmd {
	return nil
}

func (m listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case listMsg:
		listItems := make([]list.Item, len(msg.items))
		for i, pr := range msg.items {
			listItems[i] = pr
		}
		m.list.SetItems(listItems)
		m.focussed = true

	case tableMsg:
		m.focussed = false

	case resetMsg:
		m.focussed = true
	}

	var cmd tea.Cmd
	if m.focussed {
		m.list, cmd = m.list.Update(msg)
	}
	return m, cmd
}

func (m listModel) View() string {
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

func (m listModel) SelectedId() uint {
	if item, ok := m.list.SelectedItem().(*Item); ok {
		return item.Id()
	}
	return 0
}
