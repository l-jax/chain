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
	headerStyle = lipgloss.NewStyle().
			Background(grey).
			Foreground(white).
			Align(lipgloss.Center).
			Padding(0, 1)
	labelStyle = lipgloss.NewStyle().
			Background(pink).
			Foreground(white).
			Padding(0, 1)
	titleStyle = lipgloss.NewStyle().
			Foreground(darkGrey).
			Bold(true)
	selectedStyle = lipgloss.NewStyle().
			Foreground(white).
			Bold(true)
	bodyStyle = lipgloss.NewStyle().
			Foreground(darkGrey)
	helpStyle = lipgloss.NewStyle().
			Foreground(darkGrey)
)
