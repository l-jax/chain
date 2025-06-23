package tui

import (
	"chain/chain"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type Detail struct {
	pr       *chain.Pr
	chain    []chain.Pr
	loaded   bool
	quitting bool
	err      error
}

func InitDetail(chain []chain.Pr, pr *chain.Pr) *Detail {
	m := Detail{chain: chain, pr: pr}

	m.loaded = true
	return &m
}

func (m Detail) Init() tea.Cmd {
	return nil
}

func (m Detail) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	return m, cmd
}

func (m Detail) View() string {
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
			bodyStyle.Padding(0, 2).Render(m.pr.Title()),
			labelStyle.Render(m.pr.State().String()),
		),
		bodyStyle.Padding(1, 0).Render(m.pr.Body()),
		m.RenderChain(),
	)
}

func (m *Detail) RenderChain() string {
	rows := make([][]string, len(m.chain))
	for i, pr := range m.chain {
		rows[i] = []string{
			strconv.FormatUint(uint64(pr.Id()), 10),
			pr.Branch(),
			pr.State().String(),
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
