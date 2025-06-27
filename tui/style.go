package tui

import (
	"github.com/charmbracelet/lipgloss"
)

/* COLORS */
var (
	purple   = lipgloss.Color("99")
	pink     = lipgloss.Color("205")
	darkGrey = lipgloss.Color("245")
	grey     = lipgloss.Color("241")
	white    = lipgloss.Color("255")
)

/* PANE */
var (
	unfocussedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(grey)
	focussedStyle = unfocussedStyle.
			BorderForeground(pink)
)

/* TEXT */
var (
	titleStyle = lipgloss.NewStyle().
			Background(purple).
			Foreground(white).
			Padding(0, 1)
	labelStyle = lipgloss.NewStyle().
			Background(grey).
			Foreground(white).
			Padding(0, 1)
	bodyStyle = lipgloss.NewStyle().
			Foreground(grey)
	helpStyle = lipgloss.NewStyle().
			Foreground(darkGrey)
)

/* TABLE */
var (
	tableHeaderStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(grey).
				BorderBottom(true).
				Bold(false)
	tableSelectedStyle = lipgloss.NewStyle().
				Foreground(white).
				Background(purple).
				Bold(false)
)
