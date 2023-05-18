package view

import (
	tea "github.com/charmbracelet/bubbletea"
)

type mainView struct {
	loginView    tea.Model
	profileView  tea.Model
	registerView tea.Model
}

func InitialModel() mainView {
	return mainView{
		loginView:    newLoginView(),
		profileView:  newProfileView(),
		registerView: newRegisterView(),
	}
}

func (m mainView) Init() tea.Cmd {
	return nil
}

func (m mainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// return m.loginView.Update(msg)
	// return m.profileView.Update(msg)
	return m.registerView.Update(msg)
}

func (m mainView) View() string {
	// return m.loginView.View()
	// return m.profileView.View()
	return m.registerView.View()
}
