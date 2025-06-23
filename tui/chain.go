package tui

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type Chain struct {
	rootLink *Link
	chain    []Link
	loaded   bool
	quitting bool
	err      error
}

func InitChain(chain []Link, rootLink *Link) *Chain {
	m := Chain{chain: chain, rootLink: rootLink}

	m.loaded = true
	return &m
}

func (m Chain) Init() tea.Cmd {
	return nil
}

func (m Chain) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	return m, cmd
}

func (m Chain) View() string {
	if m.quitting {
		return ""
	}

	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	if !m.loaded {
		return "Loading..."
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			titleStyle.Render("chain"),
			bodyStyle.Padding(0, 2).Render(m.rootLink.Title()),
			labelStyle.Render(m.rootLink.Label().String()),
		),
		bodyStyle.Padding(1, 0).Render(m.rootLink.Body()),
		m.PrepareChainTable(),
	)
}

func (m *Chain) PrepareChainTable() string {
	rows := make([][]string, len(m.chain))
	for i, link := range m.chain {
		rows[i] = []string{
			strconv.FormatUint(uint64(link.Id()), 10),
			link.Branch(),
			link.Label().String(),
		}
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(purple)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			case row%2 == 0:
				return evenRowStyle
			default:
				return oddRowStyle
			}
		}).
		Headers("id", "branch", "state").
		Rows(rows...)

	return t.Render()
}
