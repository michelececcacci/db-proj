package feed

import (
	"context"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/util"
)

type feedView struct {
	posts list.Model
	q     *queries.Queries
	ctx   *context.Context
}

func (f feedView) Init() tea.Cmd { return nil }

func (f feedView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		}
	}
	f.posts, cmd = f.posts.Update(msg)
	return f, cmd
}

func (f feedView) View() string {
	return util.ListStyle.Render(f.posts.View()) + "\n"
}

func New(ctx *context.Context, q *queries.Queries, username string) feedView {
	rawPosts, _ := q.GetFullFeed(*ctx, username)
	f := feedView{
		ctx: ctx,
		q: q,
		posts: list.New(f(rawPosts), list.NewDefaultDelegate(), 25, 25),
	}
	return f
}
