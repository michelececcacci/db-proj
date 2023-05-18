package view

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type loginView struct {
	inputsView tea.Model
}

func (l loginView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return l.inputsView.Update(msg)
}

func newLoginView() loginView {
	inputs := make([]textinput.Model, 2)
	inputs[0] = newInput("username", 20)
	inputs[1] = newInput("password", 20)
	return loginView{inputsView: newMultipleInputsView(inputs)}
}

func (l loginView) Init() tea.Cmd {
	return l.inputsView.Init()
}

func (l loginView) View() string {
	return l.inputsView.View()
}
