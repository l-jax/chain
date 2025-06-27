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

var labelColor = map[string]lipgloss.Color{
	"open":     purple,
	"merged":   grey,
	"released": darkGrey,
	"closed":   white,
	"blocked":  pink,
}

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
			Background(grey).
			Foreground(white).
			Padding(0, 1)
	labelStyle = lipgloss.NewStyle().
			Background(pink).
			Foreground(white).
			Padding(0, 1)
	bodyStyle = lipgloss.NewStyle().
			Foreground(grey)
	helpStyle = lipgloss.NewStyle().
			Foreground(darkGrey)
)
