package signup

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/util"
	"github.com/michelececcacci/db-proj/view/components"
)

type signUp struct {
	inputsView        components.MultipleInputsView
	ctx               *context.Context
	q                 *queries.Queries
	errorView         tea.Model
	locations         list.Model
	selectingLocation bool
	selectedLocation sql.NullInt32
	help tea.Model
}

func (s signUp) View() string {
	sb := strings.Builder{}
	if s.selectingLocation {
		sb.WriteString(s.locations.View())
	} else {
		sb.WriteString(s.inputsView.View())
		sb.WriteString(s.errorView.View())
	}
	sb.WriteString(s.help.View())
	return sb.String()
}

func (s signUp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	s.help, _ = s.help.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return s, tea.Quit
		}
	}
	if s.selectingLocation {
		return s.UpdateLocation(msg)
	}
	return s.UpdateNormal(msg)
}

func (s signUp) UpdateNormal(msg tea.Msg) (tea.Model, tea.Cmd) {
	var message string
	var err error
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			err = s.submitMessage()
			message = "Submission successful"
		case "ctrl+l":
			s.selectingLocation = true
		}
	}
	m, cmd := s.inputsView.Update(msg)
	iv := m.(components.MultipleInputsView)
	s.inputsView = iv
	s.errorView, _ = s.errorView.Update(util.OptionalError{Err: err, Message: message})
	return s, cmd
}

func (r signUp) UpdateLocation(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+l":
			r.selectingLocation = false
		case "enter":
			r.selectedLocation.Valid = true
			s := r.locations.SelectedItem().FilterValue()
			id, _ := strconv.ParseInt(s, 10, 32)
			r.selectedLocation.Int32 = int32(id)
			r.selectingLocation = false
		}
	}
	r.locations, cmd = r.locations.Update(msg)
	return r, cmd
}

func (s signUp) Init() tea.Cmd {
	return nil
}

func (s signUp) getCurrentUserParams() queries.InsertUserParams {
	s.inputsView.Update(nil)
	t, err := util.ParseTime(s.inputsView.Inputs[4].Value())
	return queries.InsertUserParams{
		Username:  s.inputsView.Inputs[0].Value(),
		Nome:      util.ValidNullString(s.inputsView.Inputs[2].Value()),
		Cognome:   util.ValidNullString(s.inputsView.Inputs[3].Value()),
		Domicilio: s.selectedLocation,
		Datadinascita: sql.NullTime{
			Valid: err == nil,
			Time:  t,
		},
	}
}

func New(ctx *context.Context, q *queries.Queries) signUp {
	result, err := q.GetLocations(*ctx)
	if err != nil {
		panic(err)
	}
	locations := toLocations(result, q, ctx)
	inputs := []textinput.Model{
		components.NewInput("Username", 20),
		components.NewInput("Password", 20),
		components.NewInput("Name", 20),
		components.NewInput("Surname", 20),
		components.NewInput("Birthdate", 20),
	}
	s := signUp{
		inputsView: components.NewMultipleInputsView(inputs),
		ctx:        ctx,
		q:          q,
		errorView:  components.NewErrorView(),
		locations:  list.New(locations, list.NewDefaultDelegate(), 25, 25),
		help: help{},
	}
	s.locations.Title = "Select location"
	return s
}

func (s signUp) getCurrentPasswordParams() queries.InsertPasswordParams {
	return queries.InsertPasswordParams{
		Username:        s.inputsView.Inputs[0].Value(),
		Password:        s.inputsView.Inputs[1].Value(),
		Datainserimento: time.Now().UTC(),
	}
}

func (s signUp) submitMessage() error {
	if len(s.inputsView.Inputs[1].Value()) == 0 {
		return  errors.New("empty password")
	}
	err := s.q.InsertUser(*s.ctx, s.getCurrentUserParams())
	if err != nil {
		return err
	}
	err = s.q.InsertPassword(*s.ctx, s.getCurrentPasswordParams())
	return err
}
