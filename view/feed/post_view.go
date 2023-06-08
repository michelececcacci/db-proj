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
	var cmd tea.Cmd
	p.viewport, cmd = p.viewport.Update(msg)
	return p, cmd
}

func (p postView) View() string {
	return p.viewport.View() + "\nPress ctrl+b to go back to the feed\n"
}

func newPostView(p post, size tea.Msg) postView {
	pv := postView{
		viewport: components.NewPager(p.PostTitle, p.Content),
	}
	updated, _ := pv.Update(size)
	postView := updated.(postView)
	return postView
}
