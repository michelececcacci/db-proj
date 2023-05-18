package view

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type registerView struct {
	inputs multipleInputsView
}

func (r registerView) View() string {
	return r.inputs.View()
}

func (r registerView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return r.inputs.Update(msg)
}

func (r registerView) Init() tea.Cmd {
	return nil
}

func newRegisterView() registerView {
	inputs := []textinput.Model{
		newInput("username", 20),
		newInput("passord", 20),
		newInput("Location", 20),
	}
	return registerView{inputs: newMultipleInputsView(inputs)}
}
