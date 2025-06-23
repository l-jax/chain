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

type errMsg struct {
	err error
}

type sessionState uint

const (
	activeView sessionState = iota
	chainView
)

type Model struct {
	models   []tea.Model
	focussed sessionState
	handler  *handler
	loaded   bool
	err      error
	quitting bool
}

func InitModel() (tea.Model, error) {
	m := &Model{
		focussed: activeView,
		handler:  initHandler(),
		loaded:   false,
		err:      nil,
		quitting: false,
	}

	links, err := m.handler.FetchOpen(true)
	if err != nil {
		m.err = err
		return nil, err
	}

	chain, err := m.handler.FetchChain(links[0], true)
	if err != nil {
		m.err = err
		return nil, err
	}

	m.models = make([]tea.Model, 2)
	m.models[activeView] = InitOpen(links)
	m.models[chainView] = InitChain(chain, &links[0])
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
		case key.Matches(msg, keys.Enter) && m.focussed == activeView:
			selected := m.models[activeView].(Open).list.SelectedItem().(Link)
			m.handler.FetchChain(selected, false)
			chain, err := m.handler.FetchChain(selected, false)
			if err != nil {
				m.err = err
				return m, func() tea.Msg { return errMsg{err: err} }
			}
			m.models[chainView] = InitChain(chain, &selected)
			m.focussed = chainView
		case key.Matches(msg, keys.Back):
			m.focussed = activeView
		case key.Matches(msg, keys.Quit):
			m.quitting = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.models[m.focussed], cmd = m.models[m.focussed].Update(msg)
	return m, cmd
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

	active := m.models[activeView].View()
	chain := m.models[chainView].View()

	switch m.focussed {
	default:
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			focussedStyle.Render(active),
			unfocussedStyle.Render(chain),
		)
	case chainView:
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			unfocussedStyle.Render(active),
			focussedStyle.Render(chain),
		)
	}
}

func (m *Model) SetFocussed(state sessionState) {
	if state >= sessionState(len(m.models)) {
		return
	}
	m.focussed = state
}
