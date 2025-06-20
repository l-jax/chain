package tui

import "github.com/charmbracelet/lipgloss"

var (
	listStyle = lipgloss.NewStyle().
			Padding(1, 2).
			BorderStyle(lipgloss.HiddenBorder())
	focussedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("205"))
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)
