package view

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	loginView tea.Model
}

func InitialModel() model {
	return model{loginView: newLoginView()}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return m.loginView.View()
}
