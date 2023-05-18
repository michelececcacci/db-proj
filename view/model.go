package view

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	loginView   tea.Model
	profileView tea.Model
}

func InitialModel() model {
	return model{loginView: newLoginView(), profileView: newProfileView()}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// return m.loginView.Update(msg)
	return m.profileView.Update(msg)
}

func (m model) View() string {
	// return m.loginView.View()
	return m.profileView.View()
}
