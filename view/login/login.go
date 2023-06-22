package view

import (
	"errors"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/util"
	"github.com/michelececcacci/db-proj/view/components"
)

type loginView struct {
	inputsView components.MultipleInputsView
	errorView  tea.Model
	username   *string
	model loginModel
}

type loginModel interface {
 Authenticate(arg queries.AuthenticateParams, loginDate time.Time) (bool, error)
}

func (l loginView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			isPresent, err := l.model.Authenticate(l.getCurrentAuthParams(), time.Now())
			if err != nil {
				l.errorView, _ = l.errorView.Update(util.OptionalError{Err: err})
				break
			} else {
				if isPresent {
					l.errorView, _ = l.errorView.Update(util.OptionalError{Message: "You are authenticated"})
					s := l.getCurrentAuthParams().Username
					l.username = &s
				} else {
					l.errorView, _ = l.errorView.Update(util.OptionalError{Err: errors.New("wrong username or password")})
				}
			}
		}
	}
	m, cmd := l.inputsView.Update(msg)
	iv := m.(components.MultipleInputsView)
	l.inputsView = iv
	return l, cmd
}

func New(model loginModel) loginView {
	inputs := []textinput.Model{
		components.NewInput("username", 20),
		components.NewInput("password", 20),
	}
	return loginView{
		inputsView: components.NewMultipleInputsView(inputs),
		errorView:  components.NewErrorView(),
		model: model,
	}
}

func (l loginView) Init() tea.Cmd {
	return l.inputsView.Init()
}

func (l loginView) View() string {
	sb := strings.Builder{}
	sb.WriteString(l.inputsView.View())
	sb.WriteString(l.errorView.View())
	return sb.String()
}

func (l loginView) getCurrentAuthParams() queries.AuthenticateParams {
	return queries.AuthenticateParams{
		Username: l.inputsView.Inputs[0].Value(),
		Password: l.inputsView.Inputs[1].Value(),
	}
}

func (l loginView) GetAuthenticatedUsername() (string, error) {
	if l.username == nil {
		return "", errors.New("not authenticated")
	}
	return *l.username, nil
}
