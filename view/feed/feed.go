package feed

import (
	"time"

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

type likePutter interface {
	PutLike(arg queries.PutLikeParams) error
}

type feedView struct {
	posts          list.Model
	feedGetter     feedGetter
	likePutter likePutter
	state          state
	postView       tea.Model
	lastWindowSize tea.WindowSizeMsg // needed to init the viewport
	username       *string
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
			if f.state == listState && f.username != nil && f.posts.SelectedItem() != nil {
				f.state = singlePostState
				s := f.posts.SelectedItem()
				if s != nil {
					post := s.(post)
					f.postView = newPostView(post, f.lastWindowSize)
				}
			}
		case tea.KeyCtrlB:
			if f.state == singlePostState && f.username != nil {
				f.state = listState
			}
		case tea.KeyCtrlJ:
			if f.username != nil && f.posts.SelectedItem() != nil {
				post := f.posts.SelectedItem().(post)
				f.likePutter.PutLike(queries.PutLikeParams{
					Autorecontenuto: post.Author,
					Idcontenuto:     post.id,
					Username:        *f.username,
					Likedislike:     1,
					Timestamp:       time.Now().UTC(),
				})	
			}
		}
	case util.UpdateUsername:
		f.username = msg.Username
		f.fetchPosts()
	}
	if f.username == nil {
		return f, cmd
	}
	if f.state == listState {
		f.posts, cmd = f.posts.Update(msg)
	} else {
		f.postView, cmd = f.postView.Update(msg)
	}
	return f, cmd
}

func (f feedView) View() string {
	if f.username == nil {
		return "Not logged in\n"
	}
	if f.state == listState {
		return util.ListStyle.Render(f.posts.View()) + "\n"
	}
	return f.postView.View()
}

func New(fg feedGetter, lp likePutter, username *string) feedView {
	f := feedView{
		username:   username,
		feedGetter: fg,
		likePutter: lp,
		state:      listState,
		postView:   newPostView(post{}, nil),
	}
	if username != nil {
		f.fetchPosts()
	}
	return f
}

func (f *feedView) fetchPosts() {
	var p []list.Item
	rawPosts, _ := f.feedGetter.GetFullFeed(*f.username)
	p = toPost(rawPosts)
	f.posts = list.New(p, list.NewDefaultDelegate(), 40, 25)
}
