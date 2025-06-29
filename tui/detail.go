package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	detailWidth  = 41
	detailHeight = 9
)

type detailModel struct {
	targetLabel string
	item        *Item
	viewport    *viewport.Model
	quitting    bool
	err         error
}

func newDetail(targetLabel string) detailModel {
	v := viewport.New(detailWidth, detailHeight-2)
	return detailModel{
		targetLabel: targetLabel,
		viewport:    &v,
	}
}

func (m detailModel) Init() tea.Cmd {
	return nil
}

func (m detailModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case detailMsg:
		m.item = msg.item
		str := lipgloss.NewStyle().Width(detailWidth - 2).Render(msg.item.Text())
		m.viewport.SetContent(str)
	}

	return m, nil
}

func (m detailModel) View() string {
	if m.quitting {
		return ""
	}

	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.headerView(),
		"",
		m.viewport.View(),
	)
}

func (m detailModel) headerView() string {
	if m.item == nil {
		return "\n"
	}

	var blocked string
	if m.item.Blocked() {
		str := "blocked by #" + fmt.Sprint(m.item.DependsOn())
		blocked = labelStyle.Background(labelColor["blocked"]).Render(str)
	} else {
		blocked = ""
	}

	var target string
	if m.item.HasTargetLabel() {
		target = labelStyle.Background(darkGrey).Render(strings.ToLower(m.targetLabel))
	} else {
		target = ""
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render(m.item.Title()),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			labelStyle.Background(labelColor[m.item.State()]).Render(m.item.State()),
			target,
			blocked,
		),
	)
}
