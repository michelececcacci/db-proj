package view

import (
	"context"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/view/components"
)

type loginView struct {
	inputsView components.MultipleInputsView
	ctx        *context.Context
	q          *queries.Queries
	message    string
}

func (l loginView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			isPresent, err := l.q.Authenticate(*l.ctx, l.getCurrentAuthParams())
			if err != nil {
				l.message = err.Error()
				break
			} else {
				if isPresent == 1 {
					l.message = "You are authenticated\n"
				} else {
					l.message = "Wrong username or password\n"
				}
			}
		}
	}
	m, cmd := l.inputsView.Update(msg)
	iv := m.(components.MultipleInputsView)
	l.inputsView = iv
	return l, cmd
}

func newLoginView(ctx *context.Context, q *queries.Queries) loginView {
	inputs := []textinput.Model{
		components.NewInput("username", 20),
		components.NewInput("password", 20),
	}
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
	sb := strings.Builder{}
	sb.WriteString(l.inputsView.View())
	sb.WriteString(l.message)
	return sb.String()
}

func (r loginView) getCurrentAuthParams() queries.AuthenticateParams {
	return queries.AuthenticateParams{
		Username: r.inputsView.Inputs[0].Value(),
		Password: r.inputsView.Inputs[1].Value(),
	}
}
