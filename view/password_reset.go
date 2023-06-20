package view

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/util"
	"github.com/michelececcacci/db-proj/view/components"
)

type passwordResetView struct {
	input     textinput.Model
	username  *string
	ctx       *context.Context
	q         *queries.Queries
	errorView tea.Model
}

func (p passwordResetView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var err error
	var message string
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			err = p.insertPassword()
			message = "Submission successful"
		case "ctrl+c":
			return p, tea.Quit
		default:
			message = ""
		}
	}
	p.input, cmd = p.input.Update(msg)
	p.errorView, _ = p.errorView.Update(util.OptionalError{Err: err, Message: message})
	return p, cmd
}

func (p passwordResetView) View() string {
	if p.username == nil {
		return "Not logged in\n"
	}
	var sb strings.Builder
	sb.WriteString(*p.username + "\n")
	sb.WriteString(p.input.View() + "\n")
	sb.WriteString(p.errorView.View() + "\n")
	return sb.String()
}

func (p passwordResetView) Init() tea.Cmd {
	return nil
}

func newPasswordResetView(ctx *context.Context, q *queries.Queries, username *string) passwordResetView {
	i := textinput.New()
	i.Focus()
	return passwordResetView{
		username:  username,
		input:     i,
		ctx:       ctx,
		q:         q,
		errorView: components.NewErrorView(),
	}
}

func (p passwordResetView) insertPassword() error {
	if p.username == nil {
		return errors.New("")
	}
	pastpasswords, err := p.q.GetPastPasswords(*p.ctx, *p.username)
	if err != nil {
		return err
	}
	for _, pwd := range pastpasswords {
		if pwd == p.input.Value() {
			return errors.New("password already used")
		}
	}
	err = p.q.InsertPassword(*p.ctx, queries.InsertPasswordParams{
		Username:        *p.username,
		Password:        p.input.Value(),
		Datainserimento: time.Now().UTC(),
	})
	if err != nil {
		return err
	}
	return nil
}
