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
	chainView
)

type Model struct {
	models   []tea.Model
	adaptor  *chainAdaptor
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
	m.models[chainView] = newChain()
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
			return m, tea.Batch(
				m.loadChain,
				m.loadDetail,
			)

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
	chain := m.models[chainView].View()
	help := m.help.View(keys)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			focussedStyle.Render(list),
			lipgloss.JoinVertical(
				lipgloss.Left,
				unfocussedStyle.Render(detail),
				unfocussedStyle.Render(chain),
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

func (m Model) loadChain() tea.Msg {
	item := m.models[listView].(listModel).list.SelectedItem().(*pr)

	if item == nil {
		return nil
	}

	items, err := m.adaptor.GetItemsLinkedTo(item, true)
	if err != nil {
		return errMsg{err: err}
	}
	return chainMsg{items: items}
}

func (m Model) loadDetail() tea.Msg {
	item := m.models[listView].(listModel).list.SelectedItem().(*pr)

	if item == nil {
		return nil
	}

	return detailMsg{item: item}
}
