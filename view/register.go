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
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			// queries.InsertUserParams(ctx, )
		}
	}
	return r.inputs.Update(msg)
}

func (r registerView) Init() tea.Cmd {
	return nil
}

func newRegisterView() registerView {
	inputs := []textinput.Model{
		newInput("Username", 20),
		newInput("Password", 20),
		newInput("Name", 20),
		newInput("Surname", 20),
		newInput("Location", 20),
		newInput("Birthdate", 10),
	}
	return registerView{
		inputs: newMultipleInputsView(inputs),
	}
}
