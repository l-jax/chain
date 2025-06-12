package tui

import tea "github.com/charmbracelet/bubbletea"

func GetProgram() *tea.Program {
	m := newModel()
	return tea.NewProgram(m, tea.WithAltScreen())
}
