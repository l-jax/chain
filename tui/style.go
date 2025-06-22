package tui

import "github.com/charmbracelet/lipgloss"

var (
	purple = lipgloss.Color("99")
	pink   = lipgloss.Color("205")
	grey   = lipgloss.Color("241")
	white  = lipgloss.Color("255")
)

var (
	unfocussedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(grey)
	focussedStyle = unfocussedStyle.
			BorderForeground(pink)
	helpStyle = lipgloss.NewStyle().
			Foreground(grey)
	titleStyle = lipgloss.NewStyle().
			Background(purple).
			Foreground(white).
			Padding(0, 1)
	labelStyle = lipgloss.NewStyle().
			Background(grey).
			Foreground(white).
			Padding(0, 1)
	bodyStyle = lipgloss.NewStyle().
			Foreground(grey).
			Padding(1, 0)
	enumeratorStyle = lipgloss.NewStyle().
			Foreground(purple).
			MarginRight(1)
	itemStyle = lipgloss.NewStyle().
			Foreground(grey).
			MarginRight(1)
)
