package view

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	loginViewState *loginView
}

func initialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return m.loginViewState.View()
}
