package tui

import "github.com/charmbracelet/bubbles/key"

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Enter, k.Back, k.Filter, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Enter},
		{k.Back, k.Filter, k.Quit},
	}
}

type keyMap struct {
	Up     key.Binding
	Down   key.Binding
	Enter  key.Binding
	Back   key.Binding
	Filter key.Binding
	Quit   key.Binding
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓", "down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter", "right"),
		key.WithHelp("→", "select"),
	),
	Back: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("←", "back"),
	),
	Filter: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "filter"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
