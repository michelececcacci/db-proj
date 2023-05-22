package view

import (
	"context"
	"database/sql"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/util"
	"github.com/michelececcacci/db-proj/view/components"
)

type registerView struct {
	inputsView components.MultipleInputsView
	ctx        *context.Context
	q          *queries.Queries
	message    string
}

func (r registerView) View() string {
	sb := strings.Builder{}
	sb.WriteString(r.inputsView.View())
	sb.WriteString(r.message)
	return sb.String()
}

func (r registerView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			err := r.q.InsertUser(*r.ctx, r.getCurrentUserParams())
			if err != nil {
				r.message = err.Error()
			} else {
				r.message = "Submission successful\n"
			}
		}
	}
	// setting inputsView.Inputs to private forces us to
	// a lot of code duplication. (we would have to basically re-implment everytime we want to use it)
	m, cmd := r.inputsView.Update(msg)
	iv := m.(components.MultipleInputsView)
	r.inputsView = iv
	return r, cmd
}

func (r registerView) Init() tea.Cmd {
	return nil
}

func (r registerView) getCurrentUserParams() queries.InsertUserParams {
	r.inputsView.Update(nil)
	t, err := util.ParseTime(r.inputsView.Inputs[5].Value())
	valid := (err == nil)
	// password := util.ValidNullString(r.inputsView.Inputs[1].Value())
	return queries.InsertUserParams{
		Username: r.inputsView.Inputs[0].Value(),
		Nome:     util.ValidNullString(r.inputsView.Inputs[2].Value()),
		Cognome:  util.ValidNullString(r.inputsView.Inputs[3].Value()),
		Domicilio: util.ValidNullString(r.inputsView.Inputs[4].Value()),
		Datadinascita: sql.NullTime{
			Valid: valid,
			Time: t,
		}, 
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
		message:    "",
	}
}
