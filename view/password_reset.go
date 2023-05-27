package view

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
)

type passwordResetView struct {
	input    textinput.Model
	username string
	ctx      *context.Context
	q        *queries.Queries
	message  string
}

func (p passwordResetView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			err := p.insertPassword()
			if err != nil {
				p.message = err.Error() + "\n"
			} else {
				p.message = "Change successful\n"
			}
		case "ctrl+c":
			return p, tea.Quit
		default: 
			p.message = ""
		}
	}
	p.input, cmd = p.input.Update(msg)
	return p, cmd
}

func (p passwordResetView) View() string {
	var sb strings.Builder
	sb.WriteString(p.username + "\n")
	sb.WriteString(p.input.View() + "\n")
	sb.WriteString(p.message + "\n")
	return sb.String()
}

func (p passwordResetView) Init() tea.Cmd {
	return nil
}

func newPasswordResetView(ctx *context.Context, q *queries.Queries, username string) passwordResetView {
	i :=   textinput.New()
	i.Focus()
	return passwordResetView{
		username: username,
		input:  i,
		ctx:      ctx,
		q:        q,
	}
}

func (p passwordResetView) insertPassword() error {
	pastpasswords, err := p.q.GetPastPasswords(*p.ctx, p.username)
	if err != nil {
		return err
	}
	for _, pwd := range pastpasswords {
		if pwd == p.input.Value() {
			return errors.New("Password already used")
		}
	}
	err = p.q.InsertPassword(*p.ctx, queries.InsertPasswordParams{
		Username:        p.username,
		Password:        p.input.Value(),
		Datainserimento: time.Now().UTC(),
	})
	if err != nil {
		return err
	}
	return nil
}
