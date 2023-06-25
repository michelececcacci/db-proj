package feed

import (
	"fmt"

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
	return p.viewport.View() // DO NOT ADD ANYTHING HERE
}

func newPostView(p post, size tea.Msg) postView {
	pv := postView{
		viewport: components.NewPager(p.PostTitle+fmt.Sprintf(" Likes: %d", p.upvotes), p.Content),
	}
	updated, _ := pv.Update(size)
	postView := updated.(postView)
	return postView
}
