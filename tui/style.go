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
	white    = lipgloss.Color("255")
)

var labelColor = map[string]lipgloss.Color{
	"draft":    grey,
	"open":     pink,
	"merged":   purple,
	"released": darkGrey,
	"closed":   darkGrey,
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
			Background(purple).
			Foreground(white).
			Align(lipgloss.Center).
			Padding(0, 1)
	labelStyle = lipgloss.NewStyle().
			Foreground(white).
			Padding(0, 1)
	listItemStyle = lipgloss.NewStyle().
			Foreground(darkGrey)
	selectedStyle = listItemStyle.
			Foreground(white)
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
