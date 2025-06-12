package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss/tree"
	"os"
	"slices"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	pulls  []Pull
	active list.Model
	chain  *tree.Tree
}

func newModel() model {
	pulls := getFakePullRequests()
	active := slices.Collect(func(yield func(Pull) bool) {
		for _, p := range pulls {
			if p.state == StateOpen {
				if !yield(p) {
					return
				}
			}
		}
	})

	t := tree.Root(".").
		Child("A", "B", "C")

	m := model{
		pulls:  pulls,
		active: list.New(pullsToListItems(active), list.NewDefaultDelegate(), 0, 0),
		chain:  t,
	}

	m.active.Title = "ACTIVE"
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.active.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.active, cmd = m.active.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top,
		docStyle.Render(m.active.View()),
		docStyle.Render(m.chain.String()))
}

func main() {
	m := newModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func pullsToListItems(pulls []Pull) []list.Item {
	items := make([]list.Item, len(pulls))
	for i := range pulls {
		items[i] = pulls[i]
	}
	return items
}
