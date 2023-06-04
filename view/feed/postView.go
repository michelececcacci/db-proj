package feed

import tea "github.com/charmbracelet/bubbletea"

type postView struct {
}

func (p postView) Init() tea.Cmd { return nil }

func (p postView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}

func (p postView) View() string {
	return "\n"
}

func newPostView(p post) postView {
	return postView{}
}
