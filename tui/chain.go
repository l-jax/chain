package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Chain struct {
	rootLinkId uint
	list       list.Model
	loaded     bool
	quitting   bool
	err        error
}

func InitChain(rootLinkId uint) *Chain {
	m := Chain{
		rootLinkId: rootLinkId,
		list:       list.New([]list.Item{}, list.NewDefaultDelegate(), windowSize.Width/divisor, windowSize.Height/divisor),
	}
	m.list.SetShowHelp(false)
	m.list.Title = "chain"

	m.list.SetItems([]list.Item{
		NewLink("Chain 1", "Description for Chain 1", 1, 1),
		NewLink("Chain 2", "Description for Chain 2", 2, 2),
		NewLink("Chain 3", "Description for Chain 3", 3, 3),
	})

	m.loaded = true
	return &m
}

func (m Chain) Init() tea.Cmd {
	return nil
}

func (m Chain) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width/divisor, msg.Height/divisor)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Chain) View() string {
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
