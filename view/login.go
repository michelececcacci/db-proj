package view

import (
	"context"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/view/components"
)

type loginView struct {
	inputsView tea.Model
	ctx        *context.Context
	q          *queries.Queries
}

func (l loginView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return l.inputsView.Update(msg)
}

func newLoginView(ctx *context.Context, q *queries.Queries) loginView {
	inputs := make([]textinput.Model, 2)
	inputs[0] = components.NewInput("username", 20)
	inputs[1] = components.NewInput("password", 20)
	return loginView{
		inputsView: components.NewMultipleInputsView(inputs),
		ctx:        ctx,
		q:          q,
	}
}

func (l loginView) Init() tea.Cmd {
	return l.inputsView.Init()
}

func (l loginView) View() string {
	return l.inputsView.View()
}
