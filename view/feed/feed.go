package feed

import (
	"context"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/util"
)

type state int

const (
	listState = iota
	singlePostState
)

type feedView struct {
	posts    list.Model
	q        *queries.Queries
	ctx      *context.Context
	state    state
	postView tea.Model
}

func (f feedView) Init() tea.Cmd { return nil }

func (f feedView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return f, tea.Quit
		}
	}
	f.posts, cmd = f.posts.Update(msg)
	if f.state == listState {

	} else {
		s := f.posts.SelectedItem()
		post := s.(post)
		f.postView = newPostView(post)
	}
	return f, cmd
}

func (f feedView) View() string {
	if f.state == listState {
		return util.ListStyle.Render(f.posts.View()) + "\n"
	}
	return "\n"
}

func New(ctx *context.Context, q *queries.Queries, username string) feedView {
	rawPosts, _ := q.GetFullFeed(*ctx, username)
	f := feedView{
		ctx:      ctx,
		q:        q,
		posts:    list.New(toPost(rawPosts), list.NewDefaultDelegate(), 40, 25),
		state:    listState,
		postView: newPostView(post{}),
	}
	return f
}
