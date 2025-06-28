package tui

import (
	"github.com/charmbracelet/lipgloss"
)

/* COLORS */
var (
	pink     = lipgloss.Color("205")
	darkGrey = lipgloss.Color("245")
	grey     = lipgloss.Color("241")
	purple   = lipgloss.Color("57")
)

var labelColor = map[string]lipgloss.Color{
	"open":     purple,
	"merged":   grey,
	"released": darkGrey,
	"closed":   purple,
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
			Foreground(purple).
			Align(lipgloss.Center).
			Padding(0, 1)
	labelStyle = lipgloss.NewStyle().
			Background(pink).
			Foreground(purple).
			Padding(0, 1)
	titleStyle = lipgloss.NewStyle().
			Foreground(darkGrey).
			Bold(true)
	selectedStyle = lipgloss.NewStyle().
			Foreground(purple).
			Bold(true)
	bodyStyle = lipgloss.NewStyle().
			Foreground(darkGrey)
	helpStyle = lipgloss.NewStyle().
			Foreground(darkGrey)
)

/* TABLE */
var ()

/* SPINNER */
var (
	spinnerStyle = lipgloss.NewStyle().
		Foreground(pink)
)
