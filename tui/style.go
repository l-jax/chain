package tui

import "github.com/charmbracelet/lipgloss"

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
)

/* TABLE */
var (
	headerStyle  = lipgloss.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
	cellStyle    = lipgloss.NewStyle().Padding(0, 1)
	oddRowStyle  = cellStyle.Foreground(darkGrey)
	evenRowStyle = cellStyle.Foreground(grey)
)
