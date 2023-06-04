package view

import tea "github.com/charmbracelet/bubbletea"

type authView struct {
}

func (a authView) Init() tea.Cmd {
	return nil
}

func (a authView) View() string {
	return ""
}

func (a authView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return a, nil
}
