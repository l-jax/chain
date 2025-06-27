package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	windowSize tea.WindowSizeMsg
	divisor    = 6
)

type listMsg struct {
	items []*Item
}

type detailMsg struct {
	item   *Item
	linked []*Item
}

type errMsg struct {
	err error
}

type view uint

const (
	listView view = iota
	detailView
	tableView
)

type Model struct {
	models   []tea.Model
	handler  *chainAdaptor
	err      error
	quitting bool
}

func InitModel() (tea.Model, error) {
	m := &Model{
		handler: initChainAdaptor(),
		models:  make([]tea.Model, 3),
	}
	m.models[listView] = NewList()
	m.models[detailView] = NewDetail()
	m.models[tableView] = NewTable()
	return m, nil
}

func (m Model) Init() tea.Cmd {
	return m.loadList
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		windowSize = msg

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Enter):
			return m, m.loadDetail

		case key.Matches(msg, keys.Quit):
			m.quitting = true
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd
	for i, model := range m.models {
		if model == nil {
			continue
		}
		var cmd tea.Cmd
		m.models[i], cmd = model.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}

	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	list := m.models[listView].View()
	detail := m.models[detailView].View()
	table := m.models[tableView].View()

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		focussedStyle.Render(list),
		lipgloss.JoinVertical(
			lipgloss.Left,
			unfocussedStyle.Render(table),
			unfocussedStyle.Render(detail),
		),
	)
}

func (m Model) loadList() tea.Msg {
	items, err := m.handler.ListItems(true)
	if err != nil {
		return errMsg{err: err}
	}
	return listMsg{items: items}
}

func (m Model) loadDetail() tea.Msg {
	item := m.models[listView].(List).list.SelectedItem().(*Item)

	if item == nil {
		return nil
	}

	linkedItems, err := m.handler.GetItemsLinkedTo(item, true)
	if err != nil {
		return errMsg{err: err}
	}
	return detailMsg{item: item, linked: linkedItems}
}
