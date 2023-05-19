package view

// Subviews implement tea.Model and the main view is
// responsible for updating them and selecting the displayed one
// according to the state changes.

import (
	tea "github.com/charmbracelet/bubbletea"
)

type mainView struct {
	loginView    tea.Model
	profileView  tea.Model
	registerView tea.Model
	feedView     tea.Model
}

func InitialModel() mainView {
	return mainView{
		loginView:    newLoginView(),
		profileView:  newProfileView(),
		registerView: newRegisterView(),
		feedView:     NewFeedView(),
	}
}

func (m mainView) Init() tea.Cmd {
	return nil
}

func (m mainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// return m.loginView.Update(msg)
	// return m.profileView.Update(msg)
	// return m.registerView.Update(msg)
	return m.feedView.Update(msg)
}

func (m mainView) View() string {
	// return m.loginView.View()
	// return m.profileView.View()
	// return m.registerView.View()
	return m.feedView.View()
}
