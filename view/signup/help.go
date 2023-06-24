package signup

import tea "github.com/charmbracelet/bubbletea"

type help struct{}

func (h help) Init() tea.Cmd {
	return nil
}

func (h help) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return h, nil
}

func (h help) View() string {
	return "Press ctrl+l to toggle location selection"
}
