package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
)

type pullDelegate struct{}

func (d pullDelegate) Height() int                             { return 1 }
func (d pullDelegate) Spacing() int                            { return 0 }
func (d pullDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d pullDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	pull, ok := listItem.(Pull)
	if !ok {
		return
	}
	
	str := fmt.Sprintf("%s - %s - %s", pull.branch, pull.title, pull.state.String())

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}
