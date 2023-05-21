package view

import (
	"context"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/view/components"
)

type multipleInputsView interface {
	View() string
	Update(tea.Msg) (tea.Model, tea.Cmd)
	GetInputValueByIndex() string
	Init() tea.Cmd
}

type registerView struct {
	inputsView    tea.Model
	ctx           *context.Context
	q             *queries.Queries
	errorOccurred bool
}

func (r registerView) View() string {
	sb := strings.Builder{}
	sb.WriteString(r.inputsView.View())
	if r.errorOccurred {
		sb.WriteString("Error occurred on insertion.\n")
	} else {
		sb.WriteString("No errors so far!\n")
	}
	s := sb.String()
	return s
}

func (r registerView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			err := r.q.InsertUser(*r.ctx, r.getCurrentUserParams())
			if err != nil {
				r.errorOccurred = true
			} else {
				r.errorOccurred = false
			}
		}
	}
	r.inputsView, cmd = r.inputsView.Update(msg)
	return r, cmd
}

func (r registerView) Init() tea.Cmd {
	return nil
}

func (r registerView) getCurrentUserParams() queries.InsertUserParams {
	r.inputsView.Update(nil)
	// name := util.ValidNullString(r.inputsView.GetInputValueByIndex(2))
	// surname := util.ValidNullString(r.inputsView.GetInputValueByIndex(3))
	return queries.InsertUserParams{
		// Username: r.inputsView.GetInputValueByIndex(0),
		// Nome:     name,
		// Cognome:  surname,
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
		inputsView:    components.NewMultipleInputsView(inputs),
		ctx:           ctx,
		q:             q,
		errorOccurred: false,
	}
}
