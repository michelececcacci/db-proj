package view

import (
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type loginView struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
}

func (l *loginView) View() string {
	var b strings.Builder
	for _, input := range l.inputs {
		b.WriteString(input.View())
	}
	return b.String()
}

func (l *loginView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (l *loginView) Init() tea.Model {
	return nil
}

func newLoginView() loginView {
	inputs := make([]textinput.Model, 2)
	inputs[0] = newInput("username", 20) 
	inputs[1] = newInput("password", 20)
	m := loginView{
		inputs: inputs,
	}
	return m
}

func newInput(placeholder string, maxLenght int) textinput.Model {
	t := textinput.New()
	t.Placeholder = placeholder 
	t.CharLimit = maxLenght
	return t
}