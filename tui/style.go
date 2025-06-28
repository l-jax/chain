package tui

import (
	"github.com/charmbracelet/lipgloss"
)

/* COLORS */
var (
	pink      = lipgloss.Color("#FF5C8A")
	darkGrey  = lipgloss.Color("#3C3C43")
	lightGrey = lipgloss.Color("#D1D5DB")
	purple    = lipgloss.Color("#7C3AED")
	white     = lipgloss.Color("#FFFFFF")
)

var labelColor = map[string]lipgloss.Color{
	"draft":    lightGrey,
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
			BorderForeground(darkGrey)
	focussedStyle = unfocussedStyle.
			BorderForeground(pink)
)

/* LIST */
var (
	listHeaderStyle = lipgloss.NewStyle().
			Background(purple).
			Foreground(white).
			Padding(0, 1)
	listTitleStyle = lipgloss.NewStyle().
			Foreground(lightGrey).
			Width(16)
	listDescStyle = lipgloss.NewStyle().
			Foreground(darkGrey)
	selectedTitleStyle = lipgloss.NewStyle().
				Foreground(white).
				Width(16)
	selectedDescStyle = lipgloss.NewStyle().
				Foreground(lightGrey)
	helpStyle = lipgloss.NewStyle().
			Foreground(darkGrey)
)

/* DETAIL */
var (
	labelStyle = lipgloss.NewStyle().
			Foreground(white).
			Padding(0, 1)
	titleStyle = lipgloss.NewStyle().
			Foreground(white)
	subtitleStyle = lipgloss.NewStyle().
			Foreground(lightGrey)
)
