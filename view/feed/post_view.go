package feed

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/view/components"
)

type postView struct {
	viewport tea.Model
}

func (p postView) Init() tea.Cmd { return nil }

func (p postView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return p.viewport.Update(msg)
}

func (p postView) View() string {
	return p.viewport.View() + "Press ctrl+b to go back to the feed\n"
}

func newPostView(p post) postView {
	pv := postView{
		viewport: components.NewPager(p.PostTitle, p.Content),
	}
	return pv
}
