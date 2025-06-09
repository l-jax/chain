package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Chain struct {
	help     help.Model
	loaded   bool
	focused  int
	groups   []group
	quitting bool
}

func NewChain() *Chain {
	help := help.New()
	help.ShowAll = true
	return &Chain{help: help, focused: 0}
}

func (c *Chain) Init() tea.Cmd {
	return nil
}

func (c *Chain) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			c.quitting = true
			return c, tea.Quit
		}
	}
	res, cmd := c.groups[c.focused].Update(msg)
	if _, ok := res.(group); ok {
		c.groups[c.focused] = res.(group)
	} else {
		return res, cmd
	}
	return c, cmd
}

func (c *Chain) View() string {
	if c.quitting {
		return ""
	}
	if !c.loaded {
		return "loading..."
	}
	board := lipgloss.JoinHorizontal(
		lipgloss.Left,
		c.groups[StateOpen].View(),
		c.groups[StateMerged].View(),
	)
	return lipgloss.JoinVertical(lipgloss.Left, board, c.help.View(keys))
}
