package feed

import (
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

type feedGetter interface {
	GetFullFeed(username string) ([]queries.GetFullFeedRow, error)
}

type feedView struct {
	posts          list.Model
	feedGetter     feedGetter
	state          state
	postView       tea.Model
	lastWindowSize tea.WindowSizeMsg // needed to init the viewport
}

func (f feedView) Init() tea.Cmd { return nil }

func (f feedView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	// this message always needs to be passed down to the post view even when it's not focused
	case tea.WindowSizeMsg:
		f.lastWindowSize = msg
		f.postView, _ = f.postView.Update(msg)
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return f, tea.Quit
		case tea.KeyEnter:
			if f.state == listState {
				f.state = singlePostState
				s := f.posts.SelectedItem()
				if s != nil {
					post := s.(post)
					f.postView = newPostView(post, f.lastWindowSize)
				}
			}
		case tea.KeyCtrlB:
			if f.state == singlePostState {
				f.state = listState
			}
		}
	}
	if f.state == listState {
		f.posts, cmd = f.posts.Update(msg)
	} else {
		f.postView, cmd = f.postView.Update(msg)
	}
	return f, cmd
}

func (f feedView) View() string {
	if f.state == listState {
		return util.ListStyle.Render(f.posts.View()) + "\n"
	}
	return f.postView.View()
}

func New(fg feedGetter, username string) feedView {
	rawPosts, _ := fg.GetFullFeed(username)
	f := feedView{
		feedGetter: fg,
		posts:      list.New(toPost(rawPosts), list.NewDefaultDelegate(), 40, 25),
		state:      listState,
		postView:   newPostView(post{}, nil),
	}
	return f
}
