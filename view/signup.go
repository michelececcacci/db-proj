package view

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/util"
	"github.com/michelececcacci/db-proj/view/components"
)

type signUp struct {
	inputsView components.MultipleInputsView
	ctx        *context.Context
	q          *queries.Queries
	errorView tea.Model
}

func (s signUp) View() string {
	sb := strings.Builder{}
	sb.WriteString(s.inputsView.View())
	sb.WriteString(s.errorView.View())
	return sb.String()
}

func (r signUp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var message string
	var err error
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if len(r.inputsView.Inputs[1].Value()) == 0 {
				err =  errors.New("Empty password")
				break
			}
			err = r.q.InsertUser(*r.ctx, r.getCurrentUserParams())
			if err != nil {
				break
			}
			err = r.q.InsertPassword(*r.ctx, r.getCurrentPasswordParams())
			if err != nil {
				break
			}
			message = "Submission successful"
		}
	}
	m, cmd := r.inputsView.Update(msg)
	iv := m.(components.MultipleInputsView)
	r.inputsView = iv
	r.errorView, _ = r.errorView.Update(util.OptionalError{Err: err, Message: message})
	return r, cmd
}

func (r signUp) Init() tea.Cmd {
	return nil
}

func (r signUp) getCurrentUserParams() queries.InsertUserParams {
	r.inputsView.Update(nil)
	t, err := util.ParseTime(r.inputsView.Inputs[5].Value())
	return queries.InsertUserParams{
		Username:  r.inputsView.Inputs[0].Value(),
		Nome:      util.ValidNullString(r.inputsView.Inputs[2].Value()),
		Cognome:   util.ValidNullString(r.inputsView.Inputs[3].Value()),
		Domicilio: sql.NullString{Valid: false},
		Datadinascita: sql.NullTime{
			Valid: err == nil,
			Time:  t,
		},
	}
}

func newRegisterView(ctx *context.Context, q *queries.Queries) signUp {
	inputs := []textinput.Model{
		components.NewInput("Username", 20),
		components.NewInput("Password", 20),
		components.NewInput("Name", 20),
		components.NewInput("Surname", 20),
		components.NewInput("Location", 20),
		components.NewInput("Birthdate", 20),
	}
	return signUp{
		inputsView: components.NewMultipleInputsView(inputs),
		ctx:        ctx,
		q:          q,
		errorView: components.NewErrorView(),
	}
}

func (r signUp) getCurrentPasswordParams() queries.InsertPasswordParams {
	return queries.InsertPasswordParams{
		Username:        r.inputsView.Inputs[0].Value(),
		Password:        r.inputsView.Inputs[1].Value(),
		Datainserimento: time.Now().UTC(),
	}
}
