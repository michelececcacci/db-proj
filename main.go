package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {

}

func initialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m model) View() string {
	return ""
}


func main() {
	p := tea.NewProgram(initialModel())
	_, err := p.Run()
	if err != nil {
		panic(err)
	}
}
