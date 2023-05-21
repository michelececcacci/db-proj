package view

import (
	"context"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/util"
	components "github.com/michelececcacci/db-proj/view/components"
)

type registerView struct {
	inputsView components.MultipleInputsView
	ctx        *context.Context
	q          *queries.Queries
}

func (r registerView) View() string {
	return r.inputsView.View()
}

func (r registerView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			err := r.q.InsertUser(*r.ctx, r.getCurrentUserParams())
			if err != nil {
				// TODO handle errors
			}
		}
	}
	return r.inputsView.Update(msg)
}

func (r registerView) Init() tea.Cmd {
	return nil
}

func (r registerView) getCurrentUserParams() queries.InsertUserParams {
	name := util.ValidNullString(r.inputsView.GetInputValueByIndex(2))
	surname := util.ValidNullString(r.inputsView.GetInputValueByIndex(3))
	return queries.InsertUserParams{
		Username: r.inputsView.GetInputValueByIndex(0),
		Nome:     name,
		Cognome:  surname,
	}
}

func newRegisterView(ctx *context.Context, q *queries.Queries) registerView {
	inputs := []textinput.Model{
		components.NewInput("Username", 20),
		components.NewInput("Password", 20),
		components.NewInput("Name", 20),
		components.NewInput("Surname", 20),
		components.NewInput("Location", 20),
		components.NewInput("Birthdate", 10),
	}
	return registerView{
		inputsView: components.NewMultipleInputsView(inputs),
		ctx:        ctx,
		q:          q,
	}
}
