package view

import (
	"context"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/michelececcacci/db-proj/queries"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type feedView struct {
	table table.Model
	q     *queries.Queries
	ctx   *context.Context
}

func (f feedView) Init() tea.Cmd { return nil }

func (f feedView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if f.table.Focused() {
				f.table.Blur()
			} else {
				f.table.Focus()
			}
		case tea.KeyCtrlC:
			return f, tea.Quit
		case tea.KeyEnter:
			// TODO Show post  to user
		}
	}
	f.table, cmd = f.table.Update(msg)
	return f, cmd
}

func (f feedView) View() string {
	return baseStyle.Render(f.table.View()) + "\n"
}

func NewFeedView(ctx *context.Context, q *queries.Queries) feedView {
	columns := []table.Column{{Title: "Title", Width: 10}, {Title: "User", Width: 10}}
	rows := []table.Row{
		{"Hi", "u1"},
		{"Hi2", "u2"},
	}
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)
	defaultStyles := table.DefaultStyles()
	defaultStyles.Header = defaultStyles.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	defaultStyles.Selected = defaultStyles.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(defaultStyles)
	f := feedView{
		table: t,
		ctx:   ctx,
		q:     q,
	}
	return f
}
