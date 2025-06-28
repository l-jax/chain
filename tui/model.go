package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type view uint

const (
	listView view = iota
	detailView
	tableView
)

type Model struct {
	models   []tea.Model
	focussed view
	adaptor  *adaptor
	help     help.Model
	err      error
	quitting bool
}

func InitModel() (tea.Model, error) {
	m := &Model{
		adaptor: newAdaptor(),
		models:  make([]tea.Model, 3),
		help:    help.New(),
	}
	m.models[listView] = newList()
	m.models[detailView] = newDetail()
	m.models[tableView] = newTable()
	m.focussed = listView
	return m, nil
}

func (m Model) Init() tea.Cmd {
	return m.loadList
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Enter):
			if m.focussed == listView {
				m.focussed = tableView
				return m, tea.Batch(
					func() tea.Msg {
						return tableLoadMsg{}
					},
					m.loadTable,
					m.loadDetail,
				)
			}
		case key.Matches(msg, keys.Back):
			if m.focussed == tableView {
				m.focussed = listView
				return m, func() tea.Msg {
					return resetMsg{}
				}
			}

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
	help := m.help.View(keys)

	if m.focussed == listView {
		list = focussedStyle.Render(list)
		detail = unfocussedStyle.Render(detail)
		table = unfocussedStyle.Render(table)
	} else {
		list = unfocussedStyle.Render(list)
		detail = unfocussedStyle.Render(detail)
		table = focussedStyle.Render(table)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			list,
			lipgloss.JoinVertical(
				lipgloss.Left,
				table,
				detail,
			),
		),
		helpStyle.Render(help),
	)
}

func (m Model) loadList() tea.Msg {
	items, err := m.adaptor.ListItems(true)
	if err != nil {
		return errMsg{err: err}
	}
	return listMsg{items: items}
}

func (m Model) loadTable() tea.Msg {
	if m.models[listView].(listModel).list.SelectedItem() == nil {
		return nil
	}

	item := m.models[listView].(listModel).list.SelectedItem().(*Item)

	if item == nil {
		return nil
	}

	items, err := m.adaptor.GetItemsLinkedTo(item, true)
	if err != nil {
		return errMsg{err: err}
	}
	return tableMsg{items: items}
}

func (m Model) loadDetail() tea.Msg {
	if m.models[listView].(listModel).list.SelectedItem() == nil {
		return nil
	}

	item := m.models[listView].(listModel).list.SelectedItem().(*Item)

	if item == nil {
		return nil
	}

	return detailMsg{item: item}
}
