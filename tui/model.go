package tui

import (
	"chain/chain"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	windowSize tea.WindowSizeMsg
	divisor    = 6
)

type errMsg struct {
	err error
}

type view uint

const (
	listView view = iota
	detailView
)

type Model struct {
	models   []tea.Model
	handler  *chainAdaptor
	loaded   bool
	err      error
	quitting bool
}

func InitModel() (tea.Model, error) {
	m := &Model{
		handler:  initChainAdaptor(),
		loaded:   false,
		err:      nil,
		quitting: false,
	}

	links, err := m.handler.ListOpenPrs(true)
	if err != nil {
		m.err = err
		return nil, err
	}

	chain, err := m.handler.GetPrsLinkedTo(links[0], true)
	if err != nil {
		m.err = err
		return nil, err
	}

	m.models = make([]tea.Model, 2)
	m.models[listView] = InitList(links)
	m.models[detailView] = InitDetail(chain, links[0])
	return m, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		windowSize = msg
		if m.loaded {
			break
		}
		for i := range m.models {
			m.models[i].Update(msg)
		}
		m.loaded = true
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Enter):
			selected := m.models[listView].(List).list.SelectedItem().(*chain.Pr)
			m.handler.GetPrsLinkedTo(selected, false)
			chain, err := m.handler.GetPrsLinkedTo(selected, false)
			if err != nil {
				m.err = err
				return m, func() tea.Msg { return errMsg{err: err} }
			}
			m.models[detailView] = InitDetail(chain, selected)
		case key.Matches(msg, keys.Quit):
			m.quitting = true
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd
	for i, model := range m.models {
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

	if !m.loaded {
		return "Loading..."
	}

	list := m.models[listView].View()
	detail := m.models[detailView].View()

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		focussedStyle.Render(list),
		unfocussedStyle.Render(detail),
	)
}
